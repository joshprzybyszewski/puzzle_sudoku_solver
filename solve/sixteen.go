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
	s, ok := solveSixteen(&puzzle)
	if !ok {
		return model.Sixteen{}, errors.New(`did not solve`)
	}
	return s, nil
}

func solveSixteen(
	s *model.Sixteen,
) (model.Sixteen, bool) {

	r, c := s.Best()
	if r > s.Size() {
		if s.IsSolved() {
			return *s, true
		}
		return model.Sixteen{}, false
	}
	for i := uint8(1); i <= s.Size(); i++ {
		cpy, ok := s.Place(r, c, i)
		if !ok {
			continue
		}
		solved, ok := solveSixteen(&cpy)
		if ok {
			return solved, true
		}
	}

	return model.Sixteen{}, false
}
