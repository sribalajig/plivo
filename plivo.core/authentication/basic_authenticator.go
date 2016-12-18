package authentication

type BasicAuthenticator struct {
}

func (basicAuthenticator BasicAuthenticator) Authenticate(credentials Credentials) bool {
	return true
}

func NewBasicAuthenticator() BasicAuthenticator {
	return BasicAuthenticator{}
}
