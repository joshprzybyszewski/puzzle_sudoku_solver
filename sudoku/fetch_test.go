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
	assert.NotEmpty(t, s)
}
