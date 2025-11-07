package twitty

import (
	"context"
	"fmt"
	"regexp"
	"strings"
)

var (
	UsernameMinimumLength = 3
	PasswordMinimumLength = 6
)

var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

type AuthService interface {
	Register(ctx context.Context, input RegisterInput) (AuthResponse, error)
}

type AuthResponse struct {
	AccessToken string
	User        User
}

type RegisterInput struct {
	Username        string
	Email           string
	Password        string
	ConfirmPassword string
}

func (input *RegisterInput) Sanitize() {
	input.Email = strings.TrimSpace(input.Email)
	input.Email = strings.ToLower(input.Email)

	input.Username = strings.TrimSpace(input.Username)
}

func (input RegisterInput) Validate() error {
	if len(input.Username) < UsernameMinimumLength {
		return fmt.Errorf("%w: username not long enough, (%d) character at least", ValidationError, UsernameMinimumLength)
	}

	if len(input.Password) < PasswordMinimumLength {
		return fmt.Errorf("%w: password not long enough, (%d) character at least", ValidationError, UsernameMinimumLength)
	}

	if input.Password != input.ConfirmPassword {
		return fmt.Errorf("%w confirm password must match the password", ValidationError)
	}

	if !emailRegex.MatchString(input.Email) {
		return fmt.Errorf("%w email validation error", ValidationError)
	}

	return nil
}
