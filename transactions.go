package uphold

import "fmt"

// TransactionService works with Transaction API
type TransactionService struct {
	client *Client
}

// Create a new transaction on provided quote
func (t *TransactionService) Create(card Card, q Quote) (*Txn, *Response, error) {
	//fmt.Println("Card ID ", card.ID)
	//card.ID = "a6d35fcd-xxxx-9c9d1dda6d57"
	rel := fmt.Sprintf("me/cards/%s/transactions", card.ID)
	if q.Realtime {
		rel = rel + "?commit=true"
	}

	req, err := t.client.NewRequest("POST", rel, q)

	if err != nil {
		return nil, nil, err
	}

	txn := new(Txn)
	resp, err := t.client.Do(req, txn)

	if err != nil {
		return nil, resp, err
	}

	return txn, resp, nil
}

// Commit a pending transaction on card
func (t *TransactionService) Commit(card Card, txn Txn, msg string) (*Txn, *Response, error) {
	rel := fmt.Sprintf("me/cards/%s/transactions/%s/commit", card.ID, txn.ID)

	payload := map[string]string{"message": msg}

	req, err := t.client.NewRequest("POST", rel, payload)
	if err != nil {
		return nil, nil, err
	}

	r := new(Txn)
	resp, err := t.client.Do(req, r)
	if err != nil {
		return nil, resp, err
	}

	return r, resp, nil
}

// Cancel an unclaimed transaction on a card
func (t *TransactionService) Cancel(card Card, txn Txn) (*Txn, *Response, error) {
	rel := fmt.Sprintf("me/cards/%s/transactions/%s/cancel", card.ID, txn.ID)
	req, err := t.client.NewRequest("POST", rel, nil)
	if err != nil {
		return nil, nil, err
	}

	r := new(Txn)
	resp, err := t.client.Do(req, r)
	if err != nil {
		return nil, resp, err
	}

	return r, resp, nil
}

// Resend a reminder on an unclaimed transaction
func (t *TransactionService) Resend(card Card, txn Txn) (*Txn, *Response, error) {
	rel := fmt.Sprintf("me/cards/%s/transactions/%s/resend", card.ID, txn.ID)
	req, err := t.client.NewRequest("POST", rel, nil)
	if err != nil {
		return nil, nil, err
	}

	r := new(Txn)
	resp, err := t.client.Do(req, r)
	if err != nil {
		return nil, resp, err
	}

	return r, resp, nil
}

// ListForUser lists all the transactions for current user
func (t *TransactionService) ListForUser() (*[]Txn, *Response, error) {
	req, err := t.client.NewRequest("GET", "me/transactions", nil)
	if err != nil {
		return nil, nil, err
	}

	r := new([]Txn)
	resp, err := t.client.Do(req, r)
	if err != nil {
		return nil, resp, err
	}

	return r, resp, nil
}

// ListForCard lists all the transactions for a card
func (t *TransactionService) ListForCard(card Card) (*[]Txn, *Response, error) {
	rel := fmt.Sprintf("me/cards/%s/transactions", card.ID)
	req, err := t.client.NewRequest("GET", rel, nil)
	if err != nil {
		return nil, nil, err
	}

	r := new([]Txn)
	resp, err := t.client.Do(req, r)
	if err != nil {
		return nil, resp, err
	}

	return r, resp, nil
}

func (t *TransactionService) GetAllTransactions() (*[]Txn, *Response, error) {
	req, err := t.client.NewRequest("GET", "reserve/transactions", nil)
	if err != nil {
		return nil, nil, err
	}

	r := new([]Txn)
	resp, err := t.client.Do(req, r)
	if err != nil {
		return nil, resp, err
	}

	return r, resp, nil
}
