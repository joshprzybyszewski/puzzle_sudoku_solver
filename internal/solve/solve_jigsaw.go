package solve

import (
	"context"
)

func Jigsaw(
	ctx context.Context,
	values [][]uint8,
	boxNums [][]uint8,
) ([][]uint8, error) {

	wf := newWorkforce()
	wf.start(ctx)
	output, err := wf.solve(ctx, NewPuzzleWithBoxLookups(values, boxNums))
	wf.stop()
	if err != nil {
		return nil, err
	}

	answer := make([][]uint8, len(boxNums))
	for r := range boxNums {
		answer[r] = make([]uint8, len(boxNums[r]))
		for c := range boxNums[r] {
			answer[r][c] = uint8(output.grid[r][c])
		}
	}

	return answer, nil
}
