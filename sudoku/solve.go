package sudoku

import (
	"context"

	"github.com/joshprzybyszewski/puzzle_sudoku_solver/internal/solve"
)

func Solve(
	ctx context.Context,
	input [9][9]uint8,
) ([9][9]uint8, error) {
	out, err := solve.Classic(
		ctx,
		input,
	)
	return out, err
}

func Solve12x12(
	ctx context.Context,
	input [12][12]uint8,
) ([12][12]uint8, error) {
	out, err := solve.Twelve(
		ctx,
		input,
	)
	return out, err
}

func Solve16x16(
	ctx context.Context,
	input [16][16]uint8,
) ([16][16]uint8, error) {
	out, err := solve.Sixteen(
		ctx,
		input,
	)
	return out, err
}
