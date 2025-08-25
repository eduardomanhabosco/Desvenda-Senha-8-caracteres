package validate

import (
	"errors"
	"unicode"
)

// Valida se contem 8 DIGITOS*
func ValidateNumeric8(s string) error {
	if len(s) != 8 {
		return errors.New("a senha deve ter exatamente 8 caracteres")
	}
	for _, r := range s {
		if !unicode.IsDigit(r) {
			return errors.New("a senha deve conter apenas dígitos (0-9)")
		}
	}
	return nil
}
