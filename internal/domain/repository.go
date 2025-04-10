package domain

type AccountRepository interface {
	Save(account *Account) error
	FindById(id string) (*Account, error)
	FindByApiKey(apiKey string) (*Account, error)
	UpdateBalance(account *Account) error
}

type InvoiceRepository interface {
	Save(invoice *Invoice) error
	FindById(id string) (*Invoice, error)
	FindByAccountId(accountId string) ([]*Invoice, error)
	UpdateStatus(invoice *Invoice) error
}
