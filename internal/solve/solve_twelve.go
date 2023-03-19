package solve

import (
	"context"

	"github.com/joshprzybyszewski/puzzle_sudoku_solver/internal/model"
)

func Twelve(
	ctx context.Context,
	input model.Twelve,
) (model.Twelve, error) {
	s := make([][]uint8, len(input))
	for r := range input {
		s[r] = input[r][:]
	}

	wf := newWorkforce()
	wf.start(ctx)
	output, err := wf.solve(ctx, NewPuzzle(s))
	wf.stop()
	if err != nil {
		return model.Twelve{}, err
	}

	answer := model.Twelve{}
	for r := range answer {
		for c := range answer[r] {
			answer[r][c] = uint8(output.grid[r][c])
		}
	}

	return answer, nil
}
