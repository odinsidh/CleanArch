package entity_test

import (
	eLike "newsportal/internal/like/entity"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTargetID(t *testing.T) {
	testTable := []struct {
		name  string
		input int
		await int
		err   error
	}{
		{name: "TargetID_valid", input: 1, await: 1, err: nil},
		{name: "TargetID_zero", input: 0, await: 0, err: eLike.ErrTargetIDNotValid},
		{name: "TargetID_negative", input: -999, await: 0, err: eLike.ErrTargetIDNotValid},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			got, err := eLike.NewTargetID(testCase.input)
			require.Equal(t, testCase.await, int(got))
			require.ErrorIs(t, err, testCase.err)
		})
	}

}
