package uphold

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestTickerListAll(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/ticker", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `[{
			"ask": "1.0",
			"bid": "2.0",
			"currency": "USD",
			"pair": "USDBTC"
		},{
			"ask": "2.0",
			"bid": "3.0",
			"currency": "BTC",
			"pair": "BTCUSD"
		}]`)
	})

	want := &[]CurrencyPair{
		{Ask: 1.0, Bid: 2.0, Currency: CurrencyUSD, Pair: "USDBTC"},
		{Ask: 2.0, Bid: 3.0, Currency: CurrencyBTC, Pair: "BTCUSD"},
	}

	pairs, _, err := client.Ticker.ListAll()
	if err != nil {
		t.Fatalf("Ticker.ListAll() returned unexpected error: %v", err)
	}

	if !reflect.DeepEqual(pairs, want) {
		t.Errorf("Ticker.ListAll() returned %+v, want %+v", pairs, want)
	}
}

func TestTickerListCurrency(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/ticker/USD", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `[{
			"ask": "1.0",
			"bid": "2.0",
			"currency": "USD",
			"pair": "USDBTC"
		},{
			"ask": "2.0",
			"bid": "3.0",
			"currency": "USD",
			"pair": "BTCUSD"
		}]`)
	})

	want := &[]CurrencyPair{
		{Ask: 1.0, Bid: 2.0, Currency: CurrencyUSD, Pair: "USDBTC"},
		{Ask: 2.0, Bid: 3.0, Currency: CurrencyUSD, Pair: "BTCUSD"},
	}

	pairs, _, err := client.Ticker.List(CurrencyUSD)
	if err != nil {
		t.Fatalf("Ticker.ListAll() returned unexpected error: %v", err)
	}

	if !reflect.DeepEqual(pairs, want) {
		t.Errorf("Ticker.ListAll() returned %+v, want %+v", pairs, want)
	}
}
