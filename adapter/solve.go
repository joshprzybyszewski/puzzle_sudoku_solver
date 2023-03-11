package adapter

import (
	"time"

	smodel "github.com/joshprzybyszewski/puzzle_sudoku_solver/model"
	"github.com/joshprzybyszewski/puzzler/model"
)

func solveGame(
	g *model.Game,
	timeout time.Duration,
) error {
	iter := smodel.Iterator(g.Iterator)
	if r, c := iter.GetSize(); c == 4 {
		if r == 3 {
			return solveTwelve(g, timeout)
		}
		// panic(`TODO`)
		return solveTwelve(g, timeout)
	}

	return solveClassic(g, timeout)
}
