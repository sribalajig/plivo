package repository

import (
	"plivo/plivo.core/accounts/model"
	"plivo/plivo.persistence/postgre"
)

type AccountRepository struct {
}

func (accountRepository AccountRepository) GetPassword(userName string) (model.Account, error) {
	return postgre.GetAccount(userName), nil
}
