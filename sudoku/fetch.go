package sudoku

import (
	"context"

	"github.com/joshprzybyszewski/puzzle_sudoku_solver/internal/adapter"
	smodel "github.com/joshprzybyszewski/puzzle_sudoku_solver/internal/model"
	"github.com/joshprzybyszewski/puzzler/manual"
	"github.com/joshprzybyszewski/puzzler/model"
)

func Fetch(
	ctx context.Context,
) ([9][9]uint8, error) {
	g, err := manual.Fetch(
		ctx,
		adapter.URL,
		model.Iterator(smodel.MinIterator),
	)
	if err != nil {
		return [9][9]uint8{}, err
	}

	return adapter.ConvertClassicTask(g.Task), nil
}
