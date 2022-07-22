package authenticate

import (
	"errors"
	"github.com/atlant1da-404/artik_db/pkg/config"
)

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func New(settings *config.Settings) *Login {
	return &Login{
		Username: settings.Username,
		Password: settings.Password,
	}
}

func (l Login) Required(username, password string) error {

	if l.Username != username {
		return errors.New(authenticationErr)
	}
	if l.Password != password {
		return errors.New(authenticationErr)
	}

	return nil
}
