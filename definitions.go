package uphold

// Permission is the Uphold OAuth policy
type Permission string

// String implements Stringer interface and
// converts the Permission into a string
func (p Permission) String() string {
	return string(p)
}

// PermissionsToSlice converts a group of permissions into
// a slice of strings
func PermissionsToSlice(p []Permission) []string {
	s := []string{}
	for _, v := range p {
		s = append(s, v.String())
	}
	return s
}

// Available OAuth permissions
const (
	PermissionUserRead                        Permission = "user:read"
	PermissionCardsRead                                  = "cards:read"
	PermissionCardsWrite                                 = "cards:write"
	PermissionAccountsRead                               = "accounts:read"
	PermissionContactsRead                               = "contacts:read"
	PermissionContactsWrite                              = "contacts:write"
	PermissionTransactionsRead                           = "transactions:read"
	PermissionTransactionsDeposit                        = "transactions:deposit"
	PermissionTransactionsWithdraw                       = "transactions:withdraw"
	PermissionTransactionsTransferSelf                   = "transactions:transfer:self"
	PermissionTransactionsTransferOthers                 = "transactions:transfer:others"
	PermissionTransactionsTransferApplication            = "transactions:transfer:application"
)

// CurrencyCode is ISO 4217 code of currency
type CurrencyCode string

// String implements Stringer and converts
// currency code to a string
func (c CurrencyCode) String() string {
	return string(c)
}

// Currencies supported by Uphold
const (
	CurrencyAED CurrencyCode = "AED"
	CurrencyARS              = "ARS"
	CurrencyAUD              = "AUD"
	CurrencyBRL              = "BRL"
	CurrencyBTC              = "BTC"
	CurrencyCAD              = "CAD"
	CurrencyCHF              = "CHF"
	CurrencyCNY              = "CNY"
	CurrencyDKK              = "DKK"
	CurrencyEUR              = "EUR"
	CurrencyGBP              = "GBP"
	CurrencyHKD              = "HKD"
	CurrencyILS              = "ILS"
	CurrencyINR              = "INR"
	CurrencyJPY              = "JPY"
	CurrencyKES              = "KES"
	CurrencyMXN              = "MXN"
	CurrencyNOK              = "NOK"
	CurrencyNZD              = "NZD"
	CurrencyPHP              = "PHP"
	CurrencyPLN              = "PLN"
	CurrencySEK              = "SEK"
	CurrencySGD              = "SGD"
	CurrencyUSD              = "USD"
	CurrencyVOX              = "VOX"
	CurrencyXAG              = "XAG"
	CurrencyXAU              = "XAU"
	CurrencyXPL              = "XPL"
	CurrencyXPT              = "XPT"
)

// CurrencyDesc is the description of currencies
var CurrencyDesc = map[CurrencyCode]string{
	CurrencyAED: "United Arab Emirates Dirham",
	CurrencyARS: "Argentine Peso",
	CurrencyAUD: "Australian Dollars",
	CurrencyBRL: "Brazilian Real",
	CurrencyBTC: "Bitcoin",
	CurrencyCAD: "Canadian Dollars",
	CurrencyCHF: "Swiss Franc",
	CurrencyCNY: "Yuan",
	CurrencyDKK: "Danish Krone",
	CurrencyEUR: "Euros",
	CurrencyGBP: "Pounds",
	CurrencyHKD: "Hong Kong Dollars",
	CurrencyILS: "Israeli Sheqel",
	CurrencyINR: "Indian Rupee",
	CurrencyJPY: "Yen",
	CurrencyKES: "Kenyan Shillings",
	CurrencyMXN: "Mexican Pesos",
	CurrencyNOK: "Norwegian Krone",
	CurrencyNZD: "New Zealand Dollars",
	CurrencyPHP: "Philippine Peso",
	CurrencyPLN: "Polish Zloty",
	CurrencySEK: "Swedish Krona",
	CurrencySGD: "Singapore Dollars",
	CurrencyUSD: "U.S. dollars",
	CurrencyVOX: "Voxels",
	CurrencyXAG: "Silver",
	CurrencyXAU: "Gold",
	CurrencyXPL: "Palladium",
	CurrencyXPT: "Platinum",
}
