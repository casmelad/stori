package accounts

type AccountService struct {
	repo AccountRepository
}

func NewAccountService(repository AccountRepository) AccountService {
	return AccountService{
		repo: repository,
	}
}

func (service *AccountService) Save(transactions []AccountTransaction) error {
	err := service.repo.Save(transactions)
	if err != nil {
		return err
	}
	return nil
}

func (service *AccountService) GetAccountInfo(accountNumber string) (Account, error) {
	account, err := service.repo.GetAccountInfo(accountNumber)
	if err != nil {
		return Account{}, err
	}
	return account, nil
}
