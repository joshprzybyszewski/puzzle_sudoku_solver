package solve

import (
	"context"

	"github.com/joshprzybyszewski/puzzle_sudoku_solver/internal/model"
)

func Sixteen(
	ctx context.Context,
	puzzle model.Sixteen,
) (model.Sixteen, error) {

	wf := newWorkforce()
	wf.start(ctx)
	defer wf.stop()

	return wf.solve(ctx, puzzle)
}
