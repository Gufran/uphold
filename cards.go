package uphold

import "fmt"

// CardService works with card API endpoints
type CardService struct {
	client *Client
}

// ListAll lists all available cards
func (c *CardService) ListAll() (*[]Card, *Response, error) {
	req, err := c.client.NewRequest("GET", "me/cards", nil)
	if err != nil {
		return nil, nil, err
	}

	cards := new([]Card)
	resp, err := c.client.Do(req, cards)
	if err != nil {
		return nil, resp, err
	}

	return cards, resp, nil
}

// List the card with given ID
func (c *CardService) List(ID string) (*Card, *Response, error) {
	rel := fmt.Sprintf("me/cards/%s", ID)

	req, err := c.client.NewRequest("GET", rel, nil)
	if err != nil {
		return nil, nil, err
	}

	card := new(Card)
	resp, err := c.client.Do(req, card)
	if err != nil {
		return nil, resp, err
	}

	return card, resp, err
}

// Add a new card to user account
func (c *CardService) Add(n Card) (*Card, *Response, error) {
	payload := new(Card)
	payload.Label = n.Label
	payload.Currency = n.Currency

	req, err := c.client.NewRequest("POST", "me/cards", payload)
	if err != nil {
		return nil, nil, err
	}

	card := new(Card)
	resp, err := c.client.Do(req, card)
	if err != nil {
		return nil, resp, err
	}

	return card, resp, nil
}

// Update a card on user account
func (c *CardService) Update(o Card) (*Card, *Response, error) {
	payload := new(Card)
	payload.Label = o.Label
	if o.Settings != nil {
		payload.Settings = o.Settings
	}

	rel := fmt.Sprintf("me/cards/%s", o.ID)

	req, err := c.client.NewRequest("PATCH", rel, payload)
	if err != nil {
		return nil, nil, err
	}

	card := new(Card)
	resp, err := c.client.Do(req, card)
	if err != nil {
		return nil, resp, err
	}

	return card, resp, nil
}
