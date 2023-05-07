package auth

import "csbackend/authenticator"

type Auth struct {
	*authenticator.Authenticator
}

func New() (*Auth, error) {
	_auth, err := authenticator.New()
	if err != nil {
		return nil, err
	}

	return &Auth{
		_auth,
	}, nil
}
