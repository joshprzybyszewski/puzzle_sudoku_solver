package solve

import (
	"context"

	"github.com/joshprzybyszewski/puzzle_sudoku_solver/internal/model"
)

func Classic(
	ctx context.Context,
	s model.Classic,
) (model.Classic, error) {

	wf := newWorkforce()
	wf.start(ctx)
	output, err := wf.solve(ctx, NewPuzzleFromClassic(s))
	wf.stop()
	if err != nil {
		return model.Classic{}, err
	}

	answer := model.Classic{}
	for r := range answer {
		for c := range answer[r] {
			answer[r][c] = uint8(output.grid[r][c])
		}
	}

	return answer, nil
}
