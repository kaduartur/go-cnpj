package cnpj

import (
	"errors"
	"unicode"
)

const (
	firstPos     = 5
	secondPos    = 6
	firstDigPos  = 12
	secondDigPos = 13
	validSize    = 14
)

var (
	ErrInvalidSize             = errors.New("invalid CNPJ size")
	ErrAllDigitsEquals         = errors.New("all digits is equals")
	ErrFirstVerificationDigit  = errors.New("the first verification digit is invalid")
	ErrSecondVerificationDigit = errors.New("the second verification digit is invalid")
)

// IsValid validates the value is a valid CNPJ,
// the value can be formatted as 22.896.431/0001-10 or 22896431000110.
// If err != nil the CNPJ is valid.
func IsValid(value string) error {
	cnpj := clean(value)
	if len(cnpj) != validSize {
		return ErrInvalidSize
	}

	if isEqual(cnpj) {
		return ErrAllDigitsEquals
	}

	if cnpj[firstDigPos] != calculateDigit(cnpj[:firstDigPos], firstPos) {
		return ErrFirstVerificationDigit
	}

	if cnpj[secondDigPos] != calculateDigit(cnpj[:secondDigPos], secondPos) {
		return ErrSecondVerificationDigit
	}

	return nil
}

// calculateDigit calculates the module 11 to check
// the verification digits is valid
func calculateDigit(value []uint, pos uint) uint {
	var sum uint
	var digit uint
	for i := 0; i <= len(value)-1; i++ {
		sum += value[i] * pos
		pos--

		if pos < 2 {
			pos = 9
		}
	}

	sum %= 11
	if sum > 2 {
		digit = 11 - sum
	}

	return digit
}

// Clean cleans all invalid characters in CNPJ
// and transform all char to uint
// all numbers cannot be negative
func clean(value string) []uint {
	var cnpj []uint
	for _, v := range value {
		if unicode.IsDigit(v) {
			cnpj = append(cnpj, toUint(v))
		}
	}

	return cnpj
}

// isEqual checks if the cnpj numbers are equals
func isEqual(cnpj []uint) bool {
	for i := 1; i < len(cnpj); i++ {
		if cnpj[0] != cnpj[i] {
			return false
		}
	}

	return true
}

// toUint transforme rune in an uint value
func toUint(r rune) uint {
	return uint(r) - '0'
}
