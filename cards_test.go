package uphold

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestCardListAll(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/me/cards", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")

		fmt.Fprint(w, `[{
  "id": "123",
  "address": {
    "bitcoin": "bitcoin-addr-1"
  },
  "label": "USD card",
  "currency": "USD",
  "balance": "123.45",
  "available": "12.34",
  "addresses": [{
    "id": "addr-id-1",
    "network": "bitcoin"
  }],
  "settings": {
    "position": 1,
    "starred": true
  },
  "normalized": [{
	  "available": "123.4",
	  "balance": "567.8",
	  "currency": "USD"
  }]
}, {
  "id": "456",
  "address": {
    "bitcoin": "bitcoin-addr-2"
  },
  "label": "BTC Card #2",
  "currency": "BTC",
  "balance": "0.00",
  "available": "0.00",
  "addresses": [{
    "id": "addr-id-2",
    "network": "bitcoin"
  }, {
    "id": "addr-id-3",
    "network": "bitcoin"
  }],
  "settings": {
    "position": 7,
    "starred": true
  },
  "normalized": [{
	  "available": "123.4",
	  "balance": "567.8",
	  "currency": "USD"
  }]
}]`)

	})

	cards, _, err := client.Card.ListAll()
	if err != nil {
		t.Fatalf("Card.ListAll() returned unexpected error: %v", err)
	}

	want := &[]Card{
		{
			ID: "123",
			Address: map[string]string{
				"bitcoin": "bitcoin-addr-1",
			},
			Label:     "USD card",
			Currency:  "USD",
			Balance:   123.45,
			Available: 12.34,
			Settings:  &CardSettings{1, true},
			Addresses: &[]CardAddress{
				{
					ID:      "addr-id-1",
					Network: "bitcoin",
				},
			},
			Normalized: []NormalizedCard{
				{123.4, 567.8, "USD"},
			},
		},
		{
			ID: "456",
			Address: map[string]string{
				"bitcoin": "bitcoin-addr-2",
			},
			Label:     "BTC Card #2",
			Currency:  "BTC",
			Balance:   0.00,
			Available: 0.00,
			Settings:  &CardSettings{7, true},
			Addresses: &[]CardAddress{
				{
					ID:      "addr-id-2",
					Network: "bitcoin",
				},
				{
					ID:      "addr-id-3",
					Network: "bitcoin",
				},
			},
			Normalized: []NormalizedCard{
				{123.4, 567.8, "USD"},
			},
		},
	}

	if !reflect.DeepEqual(cards, want) {
		t.Errorf("Card.ListAll() returned %+v, want %+v", cards, want)
	}
}

func TestCardListID(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/me/cards/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")

		fmt.Fprint(w, `{
  "id": "123",
  "address": {
    "bitcoin": "bitcoin-addr-1"
  },
  "label": "USD card",
  "currency": "USD",
  "balance": "123.45",
  "available": "12.34",
  "addresses": [{
    "id": "addr-id-1",
    "network": "bitcoin"
  }],
  "settings": {
    "position": 1,
    "starred": true
  },
  "normalized": [{
	  "available": "123.4",
	  "balance": "567.8",
	  "currency": "USD"
  }]
}`)

	})

	cards, _, err := client.Card.List("1")
	if err != nil {
		t.Fatalf("Card.List(1) returned unexpected error: %v", err)
	}

	want := &Card{
		ID: "123",
		Address: map[string]string{
			"bitcoin": "bitcoin-addr-1",
		},
		Label:     "USD card",
		Currency:  "USD",
		Balance:   123.45,
		Available: 12.34,
		Settings:  &CardSettings{1, true},
		Addresses: &[]CardAddress{
			{
				ID:      "addr-id-1",
				Network: "bitcoin",
			},
		},
		Normalized: []NormalizedCard{
			{123.4, 567.8, "USD"},
		},
	}

	if !reflect.DeepEqual(cards, want) {
		t.Errorf("Card.ListAll() returned %+v, want %+v", cards, want)
	}
}

func TestCardAdd(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/me/cards", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		want := `{"currency":"USD","label":"l"}`
		testBody(t, r, want)

		fmt.Fprint(w, `{
  "id": "123",
  "address": {
    "bitcoin": "bitcoin-addr-1"
  },
  "label": "l",
  "currency": "USD",
  "balance": "123.45",
  "available": "12.34",
  "addresses": [{
    "id": "addr-id-1",
    "network": "bitcoin"
  }],
  "settings": {
    "position": 1,
    "starred": true
  },
  "normalized": [{
	  "available": "123.4",
	  "balance": "567.8",
	  "currency": "USD"
  }]
}`)

	})

	tc := Card{Label: "l", Currency: CurrencyUSD}
	cards, _, err := client.Card.Add(tc)
	if err != nil {
		t.Fatalf("Card.Add(card) returned unexpected error: %v", err)
	}

	want := &Card{
		ID: "123",
		Address: map[string]string{
			"bitcoin": "bitcoin-addr-1",
		},
		Label:     "l",
		Currency:  "USD",
		Balance:   123.45,
		Available: 12.34,
		Settings:  &CardSettings{1, true},
		Addresses: &[]CardAddress{
			{
				ID:      "addr-id-1",
				Network: "bitcoin",
			},
		},
		Normalized: []NormalizedCard{
			{123.4, 567.8, "USD"},
		},
	}

	if !reflect.DeepEqual(cards, want) {
		t.Errorf("Card.ListAll() returned %+v, want %+v", cards, want)
	}
}

func TestCardUpdate(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/me/cards/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PATCH")
		want := `{"label":"l"}`
		testBody(t, r, want)

		fmt.Fprint(w, `{
  "id": "123",
  "address": {
    "bitcoin": "bitcoin-addr-1"
  },
  "label": "l",
  "currency": "USD",
  "balance": "123.45",
  "available": "12.34",
  "addresses": [{
    "id": "addr-id-1",
    "network": "bitcoin"
  }],
  "settings": {
    "position": 1,
    "starred": true
  },
  "normalized": [{
	  "available": "123.4",
	  "balance": "567.8",
	  "currency": "USD"
  }]
}`)

	})

	tc := Card{ID: "1", Label: "l"}
	cards, _, err := client.Card.Update(tc)
	if err != nil {
		t.Fatalf("Card.Update(card) returned unexpected error: %v", err)
	}

	want := &Card{
		ID: "123",
		Address: map[string]string{
			"bitcoin": "bitcoin-addr-1",
		},
		Label:     "l",
		Currency:  "USD",
		Balance:   123.45,
		Available: 12.34,
		Settings:  &CardSettings{1, true},
		Addresses: &[]CardAddress{
			{
				ID:      "addr-id-1",
				Network: "bitcoin",
			},
		},
		Normalized: []NormalizedCard{
			{123.4, 567.8, "USD"},
		},
	}

	if !reflect.DeepEqual(cards, want) {
		t.Errorf("Card.ListAll() returned %+v, want %+v", cards, want)
	}
}
