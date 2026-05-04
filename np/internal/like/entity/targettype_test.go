package entity_test

import (
	eLike "newsportal/internal/like/entity"
	"testing"

	"github.com/stretchr/testify/require"
)

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
