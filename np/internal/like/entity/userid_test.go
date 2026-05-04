package entity_test

import (
	eLike "newsportal/internal/like/entity"
	"testing"

	"github.com/stretchr/testify/require"
)

// =====================================================================
//  UserID test
// =====================================================================

func TestUserId(t *testing.T) {
	testTable := []struct {
		name     string
		input    int
		awaiting int
		err      error
	}{
		{name: "user id zero", input: 0, awaiting: 0, err: eLike.ErrUserIDNotFound},
		{name: "user id negative", input: -1, awaiting: 0, err: eLike.ErrUserIDNotFound},
		{name: "user id valid", input: 100, awaiting: 100, err: nil},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			nUID, err := eLike.NewUserID(testCase.input)
			require.Equal(t, testCase.awaiting, int(nUID))
			require.ErrorIs(t, err, testCase.err)
		})
	}
}
