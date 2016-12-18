package repository

import (
	"plivo/plivo.core/accounts/model"
)

type AccountRepository struct {
}

func (accountRepository AccountRepository) GetPassword(userName string) (model.Account, error) {
	return model.Account{}, nil
}
