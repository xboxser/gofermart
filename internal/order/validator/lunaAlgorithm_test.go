package validator

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIsValidLuhn(t *testing.T) {

	tests := []struct {
		name     string
		number   string
		valError bool
	}{
		{name: "valid number 1", number: "79927398713", valError: false},
		{name: "valid number 2", number: "4242424242424242", valError: false},
		{name: "valid number 3", number: "378282246310005", valError: false},
		{name: "valid number 4", number: "6011000990139424", valError: false},
		{name: "valid number 5", number: "30569309025904", valError: false},
		{name: "not valid number 1", number: "79927398712", valError: true},
		{name: "not valid number 2", number: "4242424242424243", valError: true},
		{name: "not valid number 3", number: "378282246310006", valError: true},
		{name: "not valid number 4", number: "6011000990139425", valError: true},
		{name: "not valid number 5", number: "123456", valError: true},
		{name: "not valid empty", number: "", valError: true},
		{name: "not valid negative number 1", number: "-4242424242424242", valError: true},
		{name: "not valid negative number 2", number: "-1", valError: true},
		{name: "not valid zero", number: "0", valError: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			isValid := IsValidLuhn(tt.number)
			if !tt.valError {
				require.True(t, isValid)
				return
			}
			require.False(t, isValid)
		})
	}
}
