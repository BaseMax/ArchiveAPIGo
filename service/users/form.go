package users

import (
	"errors"
	"regexp"
)

type RegisterForm struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (r *RegisterForm) Validate() error {

	if len(r.Username) < 4 {
		return errors.New("your username must be at least 4 characters")
	}

	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(r.Email) {
		return errors.New("your email is not a valid email address")
	}

	if len(r.Password) < 4 {
		return errors.New("your password must be at least 4 characters")
	}

	return nil
}

type LoginForm struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
