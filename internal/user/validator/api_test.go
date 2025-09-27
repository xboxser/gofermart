package validator

import (
	"gophermart/internal/user/model"
	"gophermart/internal/validator"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestValidateModelUserAPI(t *testing.T) {
	if validator.Validate == nil {
		validator.Init()
	}

	tests := []struct {
		name     string
		user     model.APIUser
		valError bool
	}{
		{name: "valid parameters", user: model.APIUser{Login: "qwerty", Password: "pass"}, valError: false},
		{name: "empty parameters", user: model.APIUser{Login: "", Password: ""}, valError: true},
		{name: "nil password", user: model.APIUser{Login: "qwerty"}, valError: true},
		{name: "nil login", user: model.APIUser{Login: "qwerty"}, valError: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateModelUserAPI(tt.user)
			if !tt.valError {
				require.NoError(t, err)
				return
			}
			assert.Error(t, err)
		})
	}
}
