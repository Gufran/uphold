package uphold

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"sync"
	"time"

	"golang.org/x/oauth2"
)

// Various URL endpoints defined by Uphold
const (
	LiveAuthURL    = "https://uphold.com/authorize/"
	SandBoxAuthURL = "https://sandbox.uphold.com/authorize/"
	TokenAccessURL = "https://api.uphold.com/oauth2/token"
	APIURL         = "https://api.uphold.com/v0/"
)

const (
	libVersion         = "0.1"
	userAgent          = "gufran-uphold/" + libVersion
	defaultContentType = "application/json"
)

const (
	// The total number of requests possible in the current window duration
	headerRateLimit = "X-RateLimit-Limit"

	// The number of requests remaining in the current window duration
	headerRateRemaining = "X-RateLimit-Remaining"

	// The time, in UTC epoch seconds, until the end of the current window duration
	headerRateReset = "X-RateLimit-Reset"

	// The time, in seconds, until the end of the current window duration
	headerRetryAfter = "Retry-After"
)

// Credential pair to use for oAuth request
type Credential struct {
	ClientID     string
	ClientSecret string
}

// Terminals describes the endpoints on oAuth
// server and client
type Terminals struct {
	// oauth2 endpoints for authentication
	// and token request
	oauth2.Endpoint

	// URL to the service which will handle
	// third leg of the oAuth process
	RedirectURL string
}

// Client is an Uphold HTTP client
type Client struct {
	http      *http.Client
	UserAgent string

	authURL        *url.URL
	tokenAccessURL *url.URL
	apiURL         *url.URL

	rateMu sync.Mutex
	rate   RequestRate

	Ticker      *TickerService
	Account     *AccountService
	Card        *CardService
	Contact     *ContactService
	Transaction *TransactionService
}

// NewClient returns an Uphold API client
func NewClient(http *http.Client) *Client {
	authURL, _ := url.Parse(LiveAuthURL)
	tokenURL, _ := url.Parse(TokenAccessURL)
	apiURL, _ := url.Parse(APIURL)

	c := &Client{
		http:           http,
		UserAgent:      userAgent,
		authURL:        authURL,
		tokenAccessURL: tokenURL,
		apiURL:         apiURL,
		rate:           RequestRate{},
	}

	c.Ticker = &TickerService{client: c}
	c.Account = &AccountService{client: c}
	c.Card = &CardService{client: c}
	c.Contact = &ContactService{client: c}
	c.Transaction = &TransactionService{client: c}

	return c
}

// UseSandbox uses the sandbox URL for authentication
func (c *Client) UseSandbox() {
	u, _ := url.ParseRequestURI(SandBoxAuthURL)
	c.authURL = u
}

// NewRequest creates an API request. A relative URL can be provided in url,
// in which case it is resolved relative to the BaseURL of the Client.
// Relative URLs should always be specified without a preceding slash.
// If specified, the value pointed to by body is JSON encoded and included
// as the request body.
func (c *Client) NewRequest(method, urlStr string, body interface{}) (*http.Request, error) {
	rel, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	u := c.apiURL.ResolveReference(rel)

	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Accept", defaultContentType)
	if c.UserAgent != "" {
		req.Header.Add("User-Agent", c.UserAgent)
	}
	return req, nil
}

// Do sends an API request and returns the API response. The API response is
// JSON decoded and stored in the value pointed to by v, or returned as an
// error if an API error has occurred.  If v implements the io.Writer
// interface, the raw response body will be written to v, without attempting to
// first decode it.
func (c *Client) Do(req *http.Request, v interface{}) (*Response, error) {
	resp, err := c.http.Do(req)
	if err != nil {
		return nil, err
	}

	defer func() {
		// Drain up to 512 bytes and close the body to let the Transport reuse the connection
		_, _ = io.CopyN(ioutil.Discard, resp.Body, 512)
		_ = resp.Body.Close()
	}()

	response := newResponse(resp)

	c.rateMu.Lock()
	c.rate = response.RequestRate
	c.rateMu.Unlock()

	err = CheckResponse(resp)
	if err != nil {
		// even though there was an error, we still return the response
		// in case the caller wants to inspect it further
		return response, err
	}

	if v != nil {
		if w, ok := v.(io.Writer); ok {
			_, _ = io.Copy(w, resp.Body)
		} else {
			err = json.NewDecoder(resp.Body).Decode(v)
			if err == io.EOF {
				err = nil // ignore EOF errors caused by empty response body
			}
		}
	}

	return response, err
}

// CheckResponse checks the API response for errors, and returns them if
// present.  A response is considered an error if it has a status code outside
// the 200 range.  API error responses are expected to have either no response
// body, or a JSON response body that maps to ErrorResponse.  Any other
// response body will be silently ignored.
//
// The error type will be *RateLimitError for rate limit exceeded errors,
// and *TwoFactorAuthError for two-factor authentication errors.
func CheckResponse(r *http.Response) error {
	if c := r.StatusCode; 200 <= c && c <= 299 {
		return nil
	}

	errorResponse := ErrorResponse{Response: r}

	if r.StatusCode == 429 && r.Header.Get(headerRateRemaining) == "0" {
		return RateLimitError{
			ErrorResponse: errorResponse,
			RequestRate:   parseRate(r),
		}

	}
	return errorResponse
}

// RequestRate describes the API rate limit,
// available requests for current time frame,
// time at which the current frame will expire
// and number of seconds to wait before next try
type RequestRate struct {
	Limit      int
	Remaining  int
	ResetOn    time.Time
	RetryAfter int
}

// Rate specifies the current rate limit for the client as determined by the
// most recent API call.  If the client is used in a multi-user application,
// this rate may not always be up-to-date.
func (c *Client) Rate() RequestRate {
	c.rateMu.Lock()
	rate := c.rate
	c.rateMu.Unlock()
	return rate
}

// Response contains the API response and request rate information
type Response struct {
	*http.Response
	RequestRate
}

// newResponse creates a new Response for the provided http.Response.
func newResponse(r *http.Response) *Response {
	response := &Response{Response: r}
	response.RequestRate = parseRate(r)

	return response
}

// parseRate parses the rate related headers.
func parseRate(r *http.Response) RequestRate {
	var rate RequestRate
	if limit := r.Header.Get(headerRateLimit); limit != "" {
		rate.Limit, _ = strconv.Atoi(limit)
	}
	if remaining := r.Header.Get(headerRateRemaining); remaining != "" {
		rate.Remaining, _ = strconv.Atoi(remaining)
	}
	if reset := r.Header.Get(headerRateReset); reset != "" {
		if v, _ := strconv.ParseInt(reset, 10, 64); v != 0 {
			rate.ResetOn = time.Unix(v, 0)
		}
	}
	if retry := r.Header.Get(headerRetryAfter); retry != "" {
		rate.RetryAfter, _ = strconv.Atoi(retry)
	}
	return rate
}

// ConfigureOAuth returns an OAuth configuration
func ConfigureOAuth(c Credential, t Terminals, s []Permission) *oauth2.Config {
	return &oauth2.Config{
		ClientID:     c.ClientID,
		ClientSecret: c.ClientSecret,
		Endpoint:     t.Endpoint,
		RedirectURL:  t.RedirectURL,
		Scopes:       PermissionsToSlice(s),
	}
}
