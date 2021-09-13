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

func IsValid(value string) error {
	c := clean(value)
	if len(c) != validSize {
		return ErrInvalidSize
	}

	if isEqual(c) {
		return ErrAllDigitsEquals
	}

	if c[firstDigPos] != calculateDigit(c[:firstDigPos], firstPos) {
		return ErrFirstVerificationDigit
	}

	if c[secondDigPos] != calculateDigit(c[:secondDigPos], secondPos) {
		return ErrSecondVerificationDigit
	}

	return nil
}

func calculateDigit(doc []uint, pos uint) uint {
	var sum uint
	var digit uint
	for i := 0; i <= len(doc)-1; i++ {
		sum += doc[i] * pos
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

func toInt(r rune) uint {
	return uint(r) - '0'
}

func clean(c string) []uint {
	var cnpj []uint
	for _, v := range c {
		if unicode.IsDigit(v) {
			cnpj = append(cnpj, toInt(v))
		}
	}

	return cnpj
}

func isEqual(cnpj []uint) bool {
	for i := 1; i < len(cnpj); i++ {
		if cnpj[0] != cnpj[i] {
			return false
		}
	}

	return true
}
