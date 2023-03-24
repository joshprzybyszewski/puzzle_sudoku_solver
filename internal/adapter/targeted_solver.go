package adapter

import (
	"context"
	"time"

	smodel "github.com/joshprzybyszewski/puzzle_sudoku_solver/internal/model"
	"github.com/joshprzybyszewski/puzzler/model"
)

type targetedSolver struct {
	id   model.GameID
	iter model.Iterator

	timeout time.Duration
}

func NewTargetedSolver(
	iter smodel.Iterator,
	id model.GameID,
	timeout time.Duration,
) targetedSolver {
	if iter < smodel.MinIterator {
		panic(`unexpected`)
	}
	if iter > smodel.MinIterator {
		panic(`unexpected`)
	}
	if timeout < 0 {
		timeout = 0
	} else if timeout > maxTimeout {
		timeout = maxTimeout
	}

	return targetedSolver{
		iter:    model.Iterator(iter),
		id:      id,
		timeout: timeout,
	}
}

func (s targetedSolver) Iterator() model.Iterator {
	return s.iter
}

func (s targetedSolver) IteratorString() model.IteratorString {
	return model.IteratorString(smodel.Iterator(s.iter).String())
}

func (s targetedSolver) GameID() model.GameID {
	return s.id
}

func (s targetedSolver) Solve(g *model.Game) error {
	ctx, cancelFn := context.WithTimeout(context.Background(), s.timeout)
	defer cancelFn()

	return solveGame(ctx, g)
}

func (s targetedSolver) Pretty(g model.Game) string {
	return `TODO`
}
