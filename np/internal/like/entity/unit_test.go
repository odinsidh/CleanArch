package entity_test

import (
	eLike "newsportal/internal/like/entity"
	"testing"

	"github.com/stretchr/testify/require"
)

// =====================================================================
//  Reaction test
// =====================================================================

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
			},
			await: nil,
			err:   eLike.ErrReactionTypeNotExist},
		{
			name: "Reaction_bad_UserID",
			input: struct {
				ReactionType int
				UserID       int
				TargetID     int
				TargetType   string
			}{
				ReactionType: 2,
				UserID:       -100,
			},
			await: nil,
			err:   eLike.ErrUserIDNotFound},
		{
			name: "Reaction_bad_TargetID",
			input: struct {
				ReactionType int
				UserID       int
				TargetID     int
				TargetType   string
			}{
				ReactionType: 2,
				UserID:       100,
				TargetID:     -100,
			},
			await: nil,
			err:   eLike.ErrTargetIDNotValid},
		{
			name: "Reaction_bad_TargetType_Empty",
			input: struct {
				ReactionType int
				UserID       int
				TargetID     int
				TargetType   string
			}{
				ReactionType: 2,
				UserID:       100,
				TargetID:     100,
				TargetType:   "",
			},
			await: nil,
			err:   eLike.ErrTargetTypeNotFound},
		{
			name: "Reaction_bad_TargetType_NotAvaliable",
			input: struct {
				ReactionType int
				UserID       int
				TargetID     int
				TargetType   string
			}{
				ReactionType: 2,
				UserID:       100,
				TargetID:     100,
				TargetType:   "NotAvailableTagetType",
			},
			await: nil,
			err:   eLike.ErrTargetTypeNotAvailable},
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

// =====================================================================
//  ReactionType test
// =====================================================================

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

// =====================================================================
//  TargetID test
// =====================================================================

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

// =====================================================================
//  TargetType test
// =====================================================================

func TestTargetType(t *testing.T) {
	testTable := []struct {
		name  string
		input string
		await string
		err   error
	}{
		{name: "TargetType_article", input: "article", await: "article", err: nil},
		{name: "TargetType_comments", input: "comments", await: "comments", err: nil},
		{name: "TargetType_empty", input: "", await: "", err: eLike.ErrTargetTypeNotFound},
		{name: "TargetType_wrong_input", input: "sabbotagemodule", await: "", err: eLike.ErrTargetTypeNotAvailable},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			got, err := eLike.NewTargetType(testCase.input)
			require.Equal(t, testCase.await, string(got))
			require.ErrorIs(t, err, testCase.err)
		})
	}
}

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
