package adapter

import (
	"errors"
	"time"

	smodel "github.com/joshprzybyszewski/puzzle_sudoku_solver/model"
	"github.com/joshprzybyszewski/puzzler/model"
)

func solveGame(
	g *model.Game,
	timeout time.Duration,
) error {
	iter := smodel.Iterator(g.Iterator)
	if r, c := iter.GetSize(); r != 3 || c != 3 {
		return errors.New(`only classic can be solved`)
	}

	return solveClassic(g, timeout)
}
