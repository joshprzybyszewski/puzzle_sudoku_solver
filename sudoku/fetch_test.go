package sudoku

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFetch(t *testing.T) {
	ctx := context.Background()
	ctx, cancelFn := context.WithTimeout(ctx, time.Second)
	t.Cleanup(cancelFn)

	s, err := Fetch(ctx)
	require.NoError(t, err)
	for r := range s {
		for c := range s[r] {
			if s[r][c] != 0 {
				return
			}
		}
	}
	assert.Fail(t, `should not be empty but was`)
}

func TestFetch12x12(t *testing.T) {
	ctx := context.Background()
	ctx, cancelFn := context.WithTimeout(ctx, time.Second)
	t.Cleanup(cancelFn)

	s, err := Fetch12x12(ctx)
	require.NoError(t, err)
	for r := range s {
		for c := range s[r] {
			if s[r][c] != 0 {
				return
			}
		}
	}
	assert.Fail(t, `should not be empty but was`)
}

func TestFetch16x16(t *testing.T) {
	ctx := context.Background()
	ctx, cancelFn := context.WithTimeout(ctx, time.Second)
	t.Cleanup(cancelFn)

	s, err := Fetch16x16(ctx)
	require.NoError(t, err)
	for r := range s {
		for c := range s[r] {
			if s[r][c] != 0 {
				return
			}
		}
	}
	assert.Fail(t, `should not be empty but was`)
}
