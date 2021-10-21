package uphold

// TODO: Write tests for Transaction endpoint
import (
	"fmt"
	"testing"
	"time"
)

func TestTransactionCreate(t *testing.T) {
	setup()
	defer teardown()

	// mux.HandleFunc("me/cards/%s/transactions", func(w http.ResponseWriter, r *http.Request) {
	// 	testMethod(t, r, "POST")
	// 	want := `{ "denomination": { "amount": "0.1", "currency": "USD" }, "destination": "foo@bar.com" }`
	// 	testBody(t, r, want)

	// 	fmt.Fprint(w, `{
	// 		ID:      "2c326b15-7106-48be-a326-06f19e69746b",
	// 	Type:    "transfer",
	// 	Message: "",
	// 	Denomination: &Denomination{
	// 		Currency: CurrencyGBP,
	// 		Pair:     "GBPUSD",
	// 		Amount:   5.00,
	// 		Rate:     0,
	// 	},
	// 	Fees:       []string{},
	// 	Status:     "completed",
	// 	Params:     &Params{},
	// 	CreatedAt:  &time.Time{},
	// 	Normalized: []Normalized{},
	// 	Origin: Origin{
	// 		CardID:      "48ce2ac5-c038-4426-b2f8-a2bdbcc93053",
	// 		Amount:      6.56,
	// 		Base:        6.56,
	// 		Commission:  0.00,
	// 		Currency:    CurrencyUSD,
	// 		Description: "Angel Rath",
	// 		Fee:         0.00,
	// 		Rate:        1.16795,
	// 		Type:        "card",
	// 		Username:    "",
	// 		Sources: []struct {
	// 			ID     string  "json:\"id,omitempty\""
	// 			Amount float32 "json:\"amount,string,omitempty\""
	// 		}{},
	// 	},
	// 	Destination: Destination{
	// 		CardID:      "bc9b3911-4bc1-4c6d-ac05-0ae87dcfc9b",
	// 		Amount:      5.57,
	// 		Base:        5.61,
	// 		Commission:  0.04,
	// 		Currency:    CurrencyEUR,
	// 		Description: "Angel Rath",
	// 		Fees:        0,
	// 		Rate:        0.85620,
	// 		Type:        "card",
	// 	},
	//   }`)

	// })

	cardObj := &Card{
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

	quoteObj := &Quote{
		Denomination: &QuoteDenomination{

			Amount:   0.1,
			Currency: CurrencyUSD,
		},
		Origin:      "",
		Destination: "foo@bar.com",
		Realtime:    false,
	}

	_, _, err := client.Transaction.Create(*cardObj, *quoteObj)

	// txn := Txn{
	// 	ID:      "2c326b15-7106-48be-a326-06f19e69746b",
	// 	Type:    "transfer",
	// 	Message: "",
	// 	Denomination: &Denomination{
	// 		Currency: CurrencyGBP,
	// 		Pair:     "GBPUSD",
	// 		Amount:   5.00,
	// 		Rate:     0,
	// 	},
	// 	Fees:       []string{},
	// 	Status:     "completed",
	// 	Params:     &Params{},
	// 	CreatedAt:  &time.Time{},
	// 	Normalized: []Normalized{},
	// 	Origin: Origin{
	// 		CardID:      "48ce2ac5-c038-4426-b2f8-a2bdbcc93053",
	// 		Amount:      6.56,
	// 		Base:        6.56,
	// 		Commission:  0.00,
	// 		Currency:    CurrencyUSD,
	// 		Description: "Angel Rath",
	// 		Fee:         0.00,
	// 		Rate:        1.16795,
	// 		Type:        "card",
	// 		Username:    "",
	// 		Sources: []struct {
	// 			ID     string  "json:\"id,omitempty\""
	// 			Amount float32 "json:\"amount,string,omitempty\""
	// 		}{},
	// 	},
	// 	Destination: Destination{
	// 		CardID:      "bc9b3911-4bc1-4c6d-ac05-0ae87dcfc9b",
	// 		Amount:      5.57,
	// 		Base:        5.61,
	// 		Commission:  0.04,
	// 		Currency:    CurrencyEUR,
	// 		Description: "Angel Rath",
	// 		Fees:        0,
	// 		Rate:        0.85620,
	// 		Type:        "card",
	// 	},
	// }

	//want := 5.00

	// if txn.Denomination.Amount != float32(want) {
	// 	t.Errorf("Expected value 5.00 but returned %f", txn.Denomination.Amount)
	// }

	// if res.StatusCode != 200 {
	// 	t.Errorf("Expected Status Code 200 but got %d", res.StatusCode)
	// }

	if err == nil {
		fmt.Println(" Returns nil TestTransactionCreate", err)
	}

	// if !reflect.DeepEqual(txn1, txn) {
	// 	t.Errorf("Card.TransactionCreate returned %+v, want %+v", txn1, txn)
	// }

	// if err != nil {
	// 	t.Error("Error in TestTransactionCreate", err)
	// }

}

func TestTransactionCancel(t *testing.T) {
	setup()
	defer teardown()

	// mux.HandleFunc("me/cards/%s/transactions/%s/cancel", func(w http.ResponseWriter, r *http.Request) {
	// 	testMethod(t, r, "POST")

	cardObj := &Card{
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

	transObj := &Txn{
		ID:           "",
		Type:         "",
		Message:      "",
		Denomination: &Denomination{},
		Fees:         []string{},
		Status:       "",
		Params: &Params{
			Currency: CurrencyUSD,
			Margin:   0.65,
			Rate:     1.16795,
			Progress: 1,
			Pair:     "EURUSD",
			TTL:      "18000",
		},
		CreatedAt:  &time.Time{},
		Normalized: []Normalized{},
		Origin:     Origin{},
		Destination: Destination{
			CardID:      "bc9b3911-4bc1-4c6d-ac05-0ae87dcfc9b3",
			Amount:      5.57,
			Base:        5.61,
			Commission:  0.04,
			Currency:    "EUR",
			Description: "Angel Rath",
			Rate:        0,
			Type:        "card",
		},
	}

	_, _, err := client.Transaction.Cancel(*cardObj, *transObj)

	// if res.StatusCode != 200 {
	// 	t.Errorf("Expected Status Code 200 but got %d", res.StatusCode)
	// }

	if err == nil {
		fmt.Println(" Returns nil TestTransactionCreate", err)
	}
	// if err != nil {
	// 	fmt.Println("Error in TestTransactionCancel", err)
	// }
	//})

}

func TestGetAllTransactions(t *testing.T) {
	setup()
	defer teardown()
	// mux.HandleFunc("reserve/transactions", func(w http.ResponseWriter, r *http.Request) {
	// 	testMethod(t, r, "GET")
	//)}
	_, _, err := client.Transaction.GetAllTransactions()

	// if res.StatusCode != 200 {
	// 	t.Errorf("Expected Status Code 200 but got %d", res.StatusCode)
	// }

	if err == nil {
		fmt.Println(" Returns nil TestTransactionCreate", err)
	}
	// if err != nil {
	// 	fmt.Println("Error in TestTransactionCancel", err)
	// }
	//})

}
