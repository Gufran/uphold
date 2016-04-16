package uphold

import "time"

// TxnType is the type of transaction
type TxnType string

// Valid transaction types
const (
	TxnTypeTransfer   TxnType = "transfer"
	TxnTypeDeposit            = "deposit"
	TxnTypeWithdrawal         = "withdrawal"
)

// TxnStatus is the status of the transaction
type TxnStatus string

// Valid transaction statuses
const (
	TxnStatusPending   TxnStatus = "pending"
	TxnStatusWaiting             = "waiting"
	TxnStatusCancelled           = "cancelled"
	TxnStatusCompleted           = "completed"
)

// FeesType is the type of fees on transaction
type FeesType string

// Valid fees types
const (
	FeesTypeDeposit    FeesType = "deposit"
	FeesTypeExchange            = "exchange"
	FeesTypeNetwork             = "network"
	FeesTypeWithdrawal          = "withdrawal"
)

// FeesTarget is the target of fees
type FeesTarget string

// Valid fees targets
const (
	FeesTargetOrigin      FeesTarget = "origin"
	FeesTargetDestination            = "destination"
)

// DestinationType is the destination type of a transaction
type DestinationType string

// Valid destination types
const (
	DestinationTypeEmail    DestinationType = "email"
	DestinationTypeExternal                 = "external"
	DestinationTypeCard                     = "card"
)

// OriginType is the origin type of a transaction
type OriginType string

// Valid origin types
const (
	OriginTypeCard     OriginType = "card"
	OriginTypeExternal OriginType = "external"
)

// AccountStatus is the status of an account
type AccountStatus string

// Valid account types
const (
	AccountStatusOK     AccountStatus = "ok"
	AccountStatusFailed               = "failed"
)

// AccountType is the type of an account
type AccountType string

// Valid account types
const (
	AccountTypeCard AccountType = "card"
	AccountTypeSepa             = "sepa"
	AccountACH                  = "ach"
)

// Account object in Uphold
type Account struct {
	ID       string        `json:"id,omitempty"`
	Currency string        `json:"currency,omitempty"`
	Status   AccountStatus `json:"status,omitempty"`
	Label    string        `json:"label,omitempty"`
	Type     AccountType   `json:"type,omitempty"`
}

// Card object in Uphold
type Card struct {
	ID                string            `json:"id,omitempty"`
	Address           map[string]string `json:"address,omitempty"`
	Available         float32           `json:"available,string,omitempty"`
	Balance           float32           `json:"balance,string,omitempty"`
	Currency          string            `json:"currency,omitempty"`
	Label             string            `json:"label,omitempty"`
	LastTransactionAt *time.Time        `json:"lastTransactionAt,omitempty"`
	Settings          *CardSettings     `json:"settings,omitempty"`
	Addresses         *[]CardAddress    `json:"addresses,omitempty"`
	Normalized        []NormalizedCard  `json:"normalized,omitempty"`
}

// CardSettings available on card
type CardSettings struct {
	Position int  `json:"position,omitempty"`
	Starred  bool `json:"starred,omitempty"`
}

// CardAddresses available on a card
type CardAddress struct {
	ID      string `json:"id,omitempty"`
	Network string `json:"network,omitempty"`
}

// NormalizedCard information
type NormalizedCard struct {
	Available float32 `json:"available,string,omitempty"`
	Balance   float32 `json:"balance,string,omitempty"`
	Currency  string  `json:"currency,omitempty"`
}

// Contact object in Uphold
type Contact struct {
	ID        string   `json:"id,omitempty"`
	Name      string   `json:"name,omitempty"`
	FirstName string   `json:"firstName,omitempty"`
	LastName  string   `json:"lastName,omitempty"`
	Emails    []string `json:"emails,omitempty"`
	Addresses []string `json:"addresses,omitempty"`
	Company   string   `json:"company,omitempty"`
}

// CurrencyPair object in Uphold
type CurrencyPair struct {
	Ask      float32 `json:"ask,string,omitempty"`
	Bid      float32 `json:"bid,string,omitempty"`
	Currency string  `json:"currency,omitempty"`
	Pair     string  `json:"pair,omitempty"`
}

// Txn is the Transaction object in Uphold
type Txn struct {
	ID           string        `json:"id,omitempty"`
	Type         TxnType       `json:"type,omitempty"`
	Message      string        `json:"message,omitempty"`
	Denomination *Denomination `json:"denomination,omitempty"`
	Fees         []string      `json:"fees,omitempty"`
	Status       TxnStatus     `json:"status,omitempty"`
	Params       *Params       `json:"params,omitempty"`
	CreatedAt    *time.Time    `json:"createdAt,omitempty"`
	Normalized   []Normalized  `json:"normalized,omitempty"`
	Origin       Origin        `json:"origin,omitempty"`
	Destination  Destination   `json:"destination,omitempty"`
}

// Quote denotes the request to perform a transaction.
// If Realtime is set to true, the resulting transaction
// will be committed immediately.
type Quote struct {
	Denomination *QuoteDenomination `json:"denomination"`
	Origin       string             `json:"origin,omitempty"`
	Destination  string             `json:"destination,omitempty"`
	Realtime     bool               `json:"-"`
}

// QuoteDenomination is denomination available on Quote
type QuoteDenomination struct {
	Amount   float32      `json:"amount,string"`
	Currency CurrencyCode `json:"currency"`
}

// Denomination is the denomination object in Uphold
type Denomination struct {
	Currency string  `json:"currency,omitempty"`
	Pair     string  `json:"pair,omitempty"`
	Amount   float32 `json:"amount,string,omitempty"`
	Rate     float32 `json:"rate,string,omitempty"`
}

// Fees object in Uphold
type Fees struct {
	Amount   float32    `json:"amount,string,omitempty"`
	Currency string     `json:"currency,omitempty"`
	Target   FeesTarget `json:"target,omitempty"`
	Type     FeesType   `json:"type,omitempty"`
}

// Params object in uphold
type Params struct {
	Currency string  `json:"currency,omitempty"`
	Margin   float32 `json:"margin,string,omitempty"`
	Rate     float32 `json:"rate,string,omitempty"`
	Progress int     `json:"progress,omitempty"`
	Pair     string  `json:"pair,omitempty"`
	TTL      string  `json:"ttl,omitempty"`
}

// Normalized object in Uphold
type Normalized struct {
	Amount     float32 `json:"amount,string,omitempty"`
	Commission float32 `json:"commission,string,omitempty"`
	Rate       float32 `json:"rate,string,omitempty"`
	Currency   float32 `json:"currency,string,omitempty"`
}

// Origin object in Uphold
type Origin struct {
	CardID      string     `json:"Cardid,omitempty"`
	Amount      float32    `json:"amount,string,omitempty"`
	Base        float32    `json:"base,string,omitempty"`
	Commission  float32    `json:"commission,string,omitempty"`
	Currency    string     `json:"currency,omitempty"`
	Description string     `json:"description,omitempty"`
	Fee         float32    `json:"fee,string,omitempty"`
	Rate        float32    `json:"rate,string,omitempty"`
	Type        OriginType `json:"type,omitempty"`
	Username    string     `json:"username,omitempty"`

	Sources []struct {
		ID     string  `json:"id,omitempty"`
		Amount float32 `json:"amount,string,omitempty"`
	} `json:"sources,omitempty"`
}

// Destination object in Uphold
type Destination struct {
	CardID      string          `json:"CardId,omitempty"`
	Amount      float32         `json:"amount,string,omitempty"`
	Base        float32         `json:"base,string,omitempty"`
	Commission  float32         `json:"commission,string,omitempty"`
	Currency    string          `json:"currency,omitempty"`
	Description string          `json:"description,omitempty"`
	Fees        float32         `json:"fees,string,omitempty"`
	Rate        float32         `json:"rate,string,omitempty"`
	Type        DestinationType `json:"type,omitempty"`
}

// Phone object in Uphold
type Phone struct {
	ID                  string `json:"id,omitempty"`
	Verified            bool   `json:"verified,omitempty"`
	Primary             bool   `json:"primary,omitempty"`
	E164Masked          string `json:"e164Masked,omitempty"`
	NationalMasked      string `json:"nationalMasked,omitempty"`
	InternationalMasked string `json:"internationalMasked,omitempty"`
}
