package adapter

import (
	"context"
	"time"

	smodel "github.com/joshprzybyszewski/puzzle_sudoku_solver/internal/model"
	"github.com/joshprzybyszewski/puzzler/model"
)

const (
	URL = `https://www.puzzle-sudoku.com/`
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
	// min = max
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
	return URL
}

func (s solver) Solve(g *model.Game) error {
	ctx, cancelFn := context.WithTimeout(context.Background(), s.Timeout())
	defer cancelFn()

	return solveGame(ctx, g)
}

func (s solver) Pretty(g model.Game) string {
	return `TODO`
}

func (s solver) IteratorString(i model.Iterator) model.IteratorString {
	return model.IteratorString(smodel.Iterator(i).String())
}
