package authentication

type Authenticator interface {
	Authenticate(credentials Credentials) bool
}
