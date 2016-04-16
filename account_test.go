package uphold

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestAccountListAll(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/me/accounts", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `[{
  "currency": "USD",
  "id": "abcd",
  "label": "card 1",
  "status": "ok",
  "type": "card"
}, {
  "currency": "EUR",
  "id": "efgh",
  "label": "bank 1",
  "status": "ok",
  "type": "sepa"
}]`)

	})

	accounts, _, err := client.Account.ListAll()
	if err != nil {
		t.Errorf("Account.ListAll() returned unexpected error: %v", err)
	}

	want := &[]Account{
		{Currency: CurrencyUSD, ID: "abcd", Label: "card 1", Status: "ok", Type: "card"},
		{Currency: CurrencyEUR, ID: "efgh", Label: "bank 1", Status: "ok", Type: "sepa"},
	}

	if !reflect.DeepEqual(accounts, want) {
		t.Errorf("Account.ListAll() returned %+v, want %+v", accounts, want)
	}
}

func TestAccountList(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/me/accounts/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{
  "currency": "USD",
  "id": "abcd",
  "label": "card 1",
  "status": "ok",
  "type": "card"
}`)

	})

	accounts, _, err := client.Account.List("1")
	if err != nil {
		t.Errorf("Account.ListAll() returned unexpected error: %v", err)
	}

	want := &Account{Currency: CurrencyUSD, ID: "abcd", Label: "card 1", Status: "ok", Type: "card"}

	if !reflect.DeepEqual(accounts, want) {
		t.Errorf("Account.List() returned %+v, want %+v", accounts, want)
	}
}
