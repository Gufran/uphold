package uphold

import "fmt"

// TickerService works with ticker API endpoints
type TickerService struct {
	client *Client
}

// ListAll retrieces a list of tickers
func (t *TickerService) ListAll() (*[]CurrencyPair, *Response, error) {
	req, err := t.client.NewRequest("GET", "ticker", nil)
	if err != nil {
		return nil, nil, err
	}

	tickers := new([]CurrencyPair)
	resp, err := t.client.Do(req, tickers)
	if err != nil {
		return nil, resp, err
	}

	return tickers, resp, nil
}

// List ticker for a currency
func (t *TickerService) List(cur CurrencyCode) (*[]CurrencyPair, *Response, error) {
	rel := fmt.Sprintf("ticker/%s", cur)

	req, err := t.client.NewRequest("GET", rel, nil)
	if err != nil {
		return nil, nil, err
	}

	ticker := new([]CurrencyPair)
	resp, err := t.client.Do(req, ticker)
	if err != nil {
		return nil, resp, err
	}

	return ticker, resp, nil
}
