package entity_test

import (
	eLike "newsportal/internal/like/entity"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReaction(t *testing.T) {
	testTable := []struct {
		name  string
		input struct {
			ReactionType int
			UserID       int
			TargetID     int
			TargetType   string
		}
		await *eLike.Reaction
		err   error
	}{
		{
			name: "Reaction_valid",

			input: struct {
				ReactionType int
				UserID       int
				TargetID     int
				TargetType   string
			}{
				ReactionType: 2,
				UserID:       100,
				TargetID:     100,
				TargetType:   "article",
			},

			await: &eLike.Reaction{
				ReactionType: 2,
				UserID:       100,
				TargetID:     100,
				TargetType:   "article",
			},

			err: nil},

		{
			name: "Reaction_bad_ReactionType",

			input: struct {
				ReactionType int
				UserID       int
				TargetID     int
				TargetType   string
			}{
				ReactionType: 999,
				UserID:       100,
				TargetID:     100,
				TargetType:   "article",
			},

			await: nil,

			err: eLike.ErrReactionTypeNotExist},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			got, err := eLike.NewReaction(testCase.input.ReactionType,
				testCase.input.UserID,
				testCase.input.TargetID,
				testCase.input.TargetType,
			)

			if got != nil {
				require.Equal(t, testCase.await.ReactionType, got.ReactionType)
				require.Equal(t, testCase.await.UserID, got.UserID)
				require.Equal(t, testCase.await.TargetID, got.TargetID)
				require.Equal(t, testCase.await.TargetType, got.TargetType)
			}

			require.ErrorIs(t, err, testCase.err)
		})
	}

}
