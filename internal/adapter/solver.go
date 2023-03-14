package adapter

import (
	"time"

	smodel "github.com/joshprzybyszewski/puzzle_sudoku_solver/internal/model"
	"github.com/joshprzybyszewski/puzzler/model"
)

var (
	maxTimeout = 15 * time.Second
)

type solver struct {
	min model.Iterator
	max model.Iterator

	timeout time.Duration
}

func NewSolver(
	min, max smodel.Iterator,
	timeout time.Duration,
) solver {
	if min < smodel.MinIterator {
		min = smodel.MinIterator
	}
	if max > smodel.MaxIterator {
		max = smodel.MaxIterator
	}
	if timeout < 0 {
		timeout = 0
	} else if timeout > maxTimeout {
		timeout = maxTimeout
	}

	return solver{
		min:     model.Iterator(min),
		max:     model.Iterator(max),
		timeout: timeout,
	}
}

func (s solver) Min() model.Iterator {
	return s.min
}

func (s solver) Max() model.Iterator {
	return s.max
}

func (s solver) Timeout() time.Duration {
	return s.timeout
}

func (s solver) URL() string {
	return `https://www.puzzle-sudoku.com/`
}

func (s solver) Solve(g *model.Game) error {
	return solveGame(g, s.Timeout())
}

func (s solver) Pretty(g model.Game) string {
	return `TODO`
}

func (s solver) IteratorString(i model.Iterator) model.IteratorString {
	return model.IteratorString(smodel.Iterator(i).String())
}
