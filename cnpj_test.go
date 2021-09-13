package cnpj

import (
	"reflect"
	"testing"
)

func TestIsValid(t *testing.T) {
	type in struct {
		cnpj string
	}

	type out struct {
		err error
	}

	tests := []struct {
		name string
		in   in
		out  out
	}{
		{
			name: "valid cnpj with format",
			in: in{
				cnpj: "79.276.501/0001-55",
			},
		},
		{
			name: "valid cnpj without format",
			in: in{
				cnpj: "10013949000180",
			},
		},
		{
			name: "invalid cnpj format",
			in: in{
				cnpj: "1*&13949000180",
			},
			out: out{
				err: ErrInvalidSize,
			},
		},
		{
			name: "invalid first verification digit",
			in: in{
				cnpj: "10.013.949/0001-20",
			},
			out: out{
				err: ErrFirstVerificationDigit,
			},
		},
		{
			name: "invalid second verification digit",
			in: in{
				cnpj: "10.013.949/0001-81",
			},
			out: out{
				err: ErrSecondVerificationDigit,
			},
		},
		{
			name: "invalid verification digit",
			in: in{
				cnpj: "00000000000000",
			},
			out: out{
				err: ErrAllDigitsEquals,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsValid(tt.in.cnpj)
			if !reflect.DeepEqual(got, tt.out.err) {
				t.Errorf("IsValid() error = %v, wantErr %v", got, tt.out.err)
			}
		})
	}
}

func BenchmarkIsValid(b *testing.B) {
	const cnpj = "22896431000110"
	for i := 0; i < b.N; i++ {
		_ = IsValid(cnpj)
	}
}

func BenchmarkIsValidWithFormat(b *testing.B) {
	const cnpj = "22.896.431/0001-10"
	for i := 0; i < b.N; i++ {
		_ = IsValid(cnpj)
	}
}
