package model

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParametersValidate(t *testing.T) {

	tt := []struct {
		name        string
		in          Parameters
		errExpected error
	}{
		{
			name:        "Missing All params",
			in:          Parameters{},
			errExpected: ErrInvalidFilter,
		},
		{
			name: "Invalid page number",
			in: Parameters{
				Page: -1,
			},
			errExpected: ErrInvalidPage,
		},
		{
			name: "Invalid filter name",
			in: Parameters{
				Filter: "name",
			},
			errExpected: ErrInvalidFilter,
		},
		{
			name: "Positive outcome",
			in: Parameters{
				Filter: "group",
			},
			errExpected: nil,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.in.Validate()

			require.Equal(t, tc.errExpected, err)
		})
	}
}
