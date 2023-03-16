package sudoku

import (
	"context"

	"github.com/joshprzybyszewski/puzzle_sudoku_solver/internal/adapter"
	"github.com/joshprzybyszewski/puzzler/manual"
)

func Fetch(
	ctx context.Context,
) ([9][9]uint8, error) {
	g, err := manual.Fetch(
		ctx,
		`https://puzzle-sudoku.com`,
		1,
	)
	if err != nil {
		return [9][9]uint8{}, err
	}

	return adapter.ConvertClassicTask(g.Task), nil
}
