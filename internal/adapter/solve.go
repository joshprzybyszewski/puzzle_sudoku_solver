package adapter

import (
	"context"

	smodel "github.com/joshprzybyszewski/puzzle_sudoku_solver/internal/model"
	"github.com/joshprzybyszewski/puzzler/model"
)

func solveGame(
	ctx context.Context,
	g *model.Game,
) error {
	iter := smodel.Iterator(g.Iterator)
	if r, c := iter.GetSize(); c == 4 {
		if r == 3 {
			return solveTwelve(ctx, g)
		}
		return solveSixteen(ctx, g)
	}

	return solveClassic(ctx, g)
}
