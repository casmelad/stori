package accounts

type AccountRepository interface {
	Save([]AccountTransaction) error
	GetAccountInfo(string) (Account, error)
}
