package uphold

import "fmt"

// ContactService works with contact API endpoints
type ContactService struct {
	client *Client
}

// ListAll contacts for a user
func (c *ContactService) ListAll() (*[]Contact, *Response, error) {
	req, err := c.client.NewRequest("GET", "me/contacts", nil)
	if err != nil {
		return nil, nil, err
	}

	contacts := new([]Contact)
	resp, err := c.client.Do(req, contacts)
	if err != nil {
		return nil, resp, err
	}

	return contacts, resp, err
}

// List a contact by given ID
func (c *ContactService) List(ID string) (*Contact, *Response, error) {
	rel := fmt.Sprintf("me/contacts/%s", ID)

	req, err := c.client.NewRequest("GET", rel, nil)
	if err != nil {
		return nil, nil, err
	}

	contact := new(Contact)
	resp, err := c.client.Do(req, contact)
	if err != nil {
		return nil, resp, err
	}

	return contact, resp, nil
}
