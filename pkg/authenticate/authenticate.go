package authenticate

import (
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

func (l Login) Required(username, password string) bool {

	if l.Username != username {
		return false
	}
	if l.Password != password {
		return false
	}

	return true
}
