package entity_test

import (
	eLike "newsportal/internal/like/entity"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReactionType(t *testing.T) {
	testTable := []struct {
		name  string
		input int
		await int
		err   error
	}{
		{name: "ReactionType_Positive", input: 2, await: 2, err: nil},
		{name: "ReactionType_Negative", input: 3, await: 3, err: nil},
		{name: "ReactionType_NonValid", input: 999, await: 0, err: eLike.ErrReactionTypeNotExist},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			got, err := eLike.NewReactionType(testCase.input)

			require.Equal(t, testCase.await, int(got))
			require.ErrorIs(t, err, testCase.err)
		})
	}

}
