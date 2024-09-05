package utils

import (
	"errors"
	"fmt"
	"strings"
	"unicode"

	"github.com/go-playground/validator/v10"
)

type Domain string

const (
	Gmail   Domain = "gmail.com"
	Outlook Domain = "outlook.com"
	ICloud  Domain = "icloud.com"
)

var allowedDomains = map[Domain]bool{
	Gmail: true,
	Outlook: true,
	ICloud: true,
}

var (
	ErrInvalidEmailDomain = errors.New("invalid email domain. Please use either of three Gmail, Outlook or iCloud")
	ErrWeakPassword = errors.New("password must be atleast 8 characters and should contain uppercase, lowercase, special characters, number")
)

func ValidationErrorToText(e validator.FieldError) string {
	switch e.Tag() {
	case "domain_email":
		return ErrInvalidEmailDomain.Error()
	case "strong_password":
		return ErrWeakPassword.Error()
	default:
		return fmt.Sprintf("%s is not valid", e.Field())
	}
}

func ValidEmailDomain(fl validator.FieldLevel) bool {
	email := fl.Field().String()
	parts := strings.Split(email, "@")
	if len(parts) != 2 {
		return false
	}
	if !allowedDomains[Domain(parts[1])] {
		return false
	}

	return true
}

func PasswordValidator(fl validator.FieldLevel) bool {
	return isPasswordValid(fl.Field().String())
}

func isPasswordValid(password string) bool {
	var (
		hasMinLen = false
		hasUpper = false
		hasLower = false
		hasNumber = false
		hasSpecial = false
	)

	if len(password) >= 8 {
		hasMinLen = true
	}

	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsNumber(char):
			hasNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}

	return hasMinLen && hasUpper && hasLower && hasNumber && hasSpecial
}