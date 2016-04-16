### Uphold

Uphold is a Go library for accessing [Uphold API][]

[![GoDoc](https://godoc.org/github.com/gufran/uphold?status.svg)](https://godoc.org/github.com/gufran/uphold) [![Build Status](https://travis-ci.org/Gufran/uphold.svg?branch=master)](https://travis-ci.org/Gufran/uphold)

### Install

```sh
go get -u github.com/gufran/uphold
```

### Usage

```go
import( 
    "http"
    "github.com/gufran/uphold"
)

func main() {
    client := uphold.NewClient(http.DefaultClient)

    // List all cards of a user
    cards, _, err := client.Card.ListAll()
    if err != nil {
        log.Fatalf("unexpected error: %s", err)
    }

    fmt.Printf("%+v", cards)
}
```

Note that this example will not actually work since Uphold API require OAuth token. The library
does not directly handle the authorization, instead pass an oauth http client to `NewClient`
method.

### Authentication

If you have an oauth token you can use it for authentication

```go
func main() {
    tokenSource := oauth2.StaticTokenSource(
        &oauth2.Token{AccessToken: "<oauthtoken>"},
    )

    authClient := oauth2.NewClient(oauth2.NoContext, tokenSource)
    client := uphold.NewClient(authClient)
  
    cards, _, err := client.Card.ListAll()
    if err != nil {
        log.Fatalf("unexpected error: %s", err)
    }

    fmt.Printf("%+v", cards)
}
```

There is also `ConfigureOAuth` method if you wish to initiate the OAuth process to authenticate
a new user. Use it like this

```go
// This example assumes that you have a webserver running to facilitate the
// oauth steps

var (
    cred = uphold.Credential{
        ClientID: "<client id>",
        ClientSecret: "<client secret>",
    }

    terms = uphold.Terminals{
        AuthURL: "http://url.to-oauth.provider/",
        TokenURL: "http://url.to-oauth.provider/token-exchange",
        RedirectURL: "http://url.to-your-app.oauth/handler",
    }

    scopes = []uphold.Permissions{
        uphold.PermissionAccountsRead,
        uphold.PermissionUserRead,
        uphold.PermissionCardsRead,
        uphold.PermissionTransactionDeposit,
        uphold.PermissionTransferApplication,
    }
)

randomStateString := "completely random string"

// Initiate the oauth process with app
func InitiateOAuthHandler(w http.ResponseWriter, r *http.Request) {
    oauthConf := uphold.ConfigureOAuth(cred, terms, scopes)

    url := oauthConf.AuthCodeURL(randomStateString, oauth2.AccessTypeOffline)
    http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

// handle callback from oauth provider and generat token
func OAuthCallbackHandler(w http.ResponseWriter, r *http.Request) {
    if state := r.FormValue("state"); state != randomStateString {
        fmt.Print("invalid state. cannot proceed")
        // ... handle the condition, redirect user or log
        return
    }

    code := r.FormValue("code")
    token, err := oauthConf.Exchange(oauth2.NoContext, code)
    if err != nil {
        fmt.Printf("failed to exchange the code for token: %s", err)
        // ... handle the condition, redirect the user or log
        return
    }

    // ... Store the token somewhere safe to use it later

    authClient := oauth2.NewClient(oauth2.NoContext, token)
    client := uphold.NewClient(authClient)
  
    cards, _, err := client.Card.ListAll()
    if err != nil {
        log.Fatalf("unexpected error: %s", err)
    }

    fmt.Printf("%+v", cards)
}
```

### TODO

 1. Tests for Transaction service
 1. Users endpoint service
 1. Transparency endpoint service
 1. Pagination support

### Contribution

Any reasonable pull request is greatly appreciated. Please fork the repository and hack away.

### License

The MIT License (MIT)

Copyright (c) 2016 Mohammad Gufran

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.


[Uphold API]: https://uphold.com/en/developer/api/documentation/
