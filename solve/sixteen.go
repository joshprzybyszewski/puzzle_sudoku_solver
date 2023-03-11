package solve

import (
	"errors"
	"time"

	"github.com/joshprzybyszewski/puzzle_sudoku_solver/model"
)

func Sixteen(
	puzzle model.Sixteen,
	timeout time.Duration,
) (model.Sixteen, error) {
	s, ok := solveSixteen(puzzle)
	if !ok {
		return model.Sixteen{}, errors.New(`did not solve`)
	}
	return s, nil
}

func solveSixteen(
	s model.Sixteen,
) (model.Sixteen, bool) {
	for r := uint8(0); r < s.Size(); r++ {
		for c := uint8(0); c < s.Size(); c++ {
			if s.IsSet(r, c) {
				continue
			}
			for i := uint8(1); i <= s.Size(); i++ {
				cpy, ok := s.Place(r, c, i)
				if !ok {
					continue
				}
				solved, ok := solveSixteen(cpy)
				if ok {
					return solved, true
				}
			}

			return model.Sixteen{}, false
		}
	}

	return s, s.IsSolved()
}
