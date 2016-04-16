package uphold

import "fmt"

// AccountService works with Account API endpoint
type AccountService struct {
	client *Client
}

// ListAll accounts for a user
func (a *AccountService) ListAll() (*[]Account, *Response, error) {
	req, err := a.client.NewRequest("GET", "me/accounts", nil)
	if err != nil {
		return nil, nil, err
	}

	accounts := new([]Account)
	resp, err := a.client.Do(req, accounts)
	if err != nil {
		return nil, resp, err
	}

	return accounts, resp, nil
}

// List a user account owned by the app
func (a *AccountService) List(ID string) (*Account, *Response, error) {
	rel := fmt.Sprintf("me/accounts/%s", ID)

	req, err := a.client.NewRequest("GET", rel, nil)
	if err != nil {
		return nil, nil, err
	}

	account := new(Account)
	resp, err := a.client.Do(req, account)
	if err != nil {
		return nil, resp, err
	}

	return account, resp, nil
}
