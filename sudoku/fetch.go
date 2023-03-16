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

func Fetch12x12(
	ctx context.Context,
) ([12][12]uint8, error) {
	g, err := manual.Fetch(
		ctx,
		adapter.URL,
		model.Iterator(smodel.Iterator12x12),
	)
	if err != nil {
		return [12][12]uint8{}, err
	}

	return adapter.ConvertTwelveTask(g.Task), nil
}

func Fetch16x16(
	ctx context.Context,
) ([16][16]uint8, error) {
	g, err := manual.Fetch(
		ctx,
		adapter.URL,
		model.Iterator(smodel.Iterator16x16),
	)
	if err != nil {
		return [16][16]uint8{}, err
	}

	sud := smodel.NewSixteen(g.Task)
	return sud.Grid(), nil
}
