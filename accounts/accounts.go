package accounts

// Account Struct
type Account struct {
	owner   string
	balance int
}

// NewAccount creates Account
func Newaccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}
