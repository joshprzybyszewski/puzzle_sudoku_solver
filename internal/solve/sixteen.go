package solve

import (
	"context"

	"github.com/joshprzybyszewski/puzzle_sudoku_solver/internal/model"
)

func Sixteen(
	ctx context.Context,
	s model.Sixteen,
) (model.Sixteen, error) {

	wf := newWorkforce()
	wf.start(ctx)
	output, err := wf.solve(ctx, NewPuzzleFromSixteen(s))
	wf.stop()
	if err != nil {
		return model.Sixteen{}, err
	}

	answer := model.Sixteen{}
	for r := range answer {
		for c := range answer[r] {
			answer[r][c] = uint8(output.grid[r][c])
		}
	}

	return answer, nil
}
