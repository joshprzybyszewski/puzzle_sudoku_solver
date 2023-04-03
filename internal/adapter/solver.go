package adapter

import (
	"context"
	"time"

	smodel "github.com/joshprzybyszewski/puzzle_sudoku_solver/internal/model"
	"github.com/joshprzybyszewski/puzzler/model"
)

const (
	JigsawURL = `https://www.puzzle-jigsaw-sudoku.com/`
	URL       = `https://www.puzzle-sudoku.com/`
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
	if min < smodel.MinIteratorJigsaw {
		min = smodel.MinIteratorJigsaw
	}
	if max > smodel.MaxIteratorJigsaw {
		max = smodel.MaxIteratorJigsaw
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

func (s solver) CanSolve(iter model.Iterator) bool {
	i := smodel.Iterator(iter)
	if i < smodel.MinIteratorStandard || i > smodel.MaxIteratorJigsaw {
		return false
	}
	if i > smodel.MaxIteratorStandard && i < smodel.MinIteratorJigsaw {
		return false
	}
	return true
}

func (s solver) Timeout() time.Duration {
	return s.timeout
}

func (s solver) URL(i model.Iterator) string {
	if i >= model.Iterator(smodel.MinIteratorJigsaw) {
		return JigsawURL
	}
	return URL
}

func (s solver) Solve(g *model.Game) error {
	ctx, cancelFn := context.WithTimeout(context.Background(), s.Timeout())
	defer cancelFn()

	if g.Iterator >= model.Iterator(smodel.MinIteratorJigsaw) {
		return solveJigsaw(ctx, g)
	}

	return solveGame(ctx, g)
}

func (s solver) Pretty(g model.Game) string {
	/* I may want to switch on size in the future
	if r, c := smodel.Iterator(g.Iterator).GetSize(); c == 4 {
		if r == 3 {
			return string(g.Answer)
		}
		return string(g.Answer)
	}
	*/

	return string(g.ID) + `: ` + string(g.Answer)
}

func (s solver) IteratorString(i model.Iterator) model.IteratorString {
	return model.IteratorString(smodel.Iterator(i).String())
}
