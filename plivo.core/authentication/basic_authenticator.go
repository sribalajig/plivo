package authentication

import (
	"plivo/plivo.core/accounts/repository"
)

type BasicAuthenticator struct {
	accountRepository repository.AccountRepository
}

func (basicAuthenticator BasicAuthenticator) Authenticate(credentials Credentials) bool {
	account, _ := basicAuthenticator.accountRepository.GetPassword(credentials.UserName)

	return account.AuthId == credentials.Password
}

func NewBasicAuthenticator() BasicAuthenticator {
	return BasicAuthenticator{}
}
