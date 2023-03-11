package solve

import (
	"errors"
	"time"

	"github.com/joshprzybyszewski/puzzle_sudoku_solver/model"
)

func Twelve(
	puzzle model.Twelve,
	timeout time.Duration,
) (model.Twelve, error) {
	s, ok := solveTwelve(puzzle)
	if !ok {
		return model.Twelve{}, errors.New(`did not solve`)
	}
	return s, nil
}

func solveTwelve(
	s model.Twelve,
) (model.Twelve, bool) {
	for r := range s {
		for c := range s[r] {
			if s[r][c] != 0 {
				continue
			}

			for i := 1; i <= len(s); i++ {
				s[r][c] = uint8(i)
				if !isValidTwelve(s) {
					continue
				}
				solved, ok := solveTwelve(s)
				if ok {
					return solved, true
				}
			}

			return model.Twelve{}, false
		}
	}

	return s, isValidTwelve(s)
}

func isValidTwelve(
	p model.Twelve,
) bool {
	var seen, b uint16
	// check each row that it has all the numbers
	for r := 0; r < len(p); r++ {
		seen = 0
		for c := range p[r] {
			if p[r][c] == 0 {
				continue
			}
			b = 1 << p[r][c]
			if seen&b == b {
				return false
			}
			seen |= b
		}
	}

	// check each col that it has all the numbers
	for c := 0; c < len(p[0]); c++ {
		seen = 0
		for r := 0; r < len(p); r++ {
			if p[r][c] == 0 {
				continue
			}
			b = 1 << p[r][c]
			if seen&b == b {
				return false
			}
			seen |= b
		}
	}

	// check each box that it has all the numbers
	for box := 0; box < len(p[0]); box++ {
		seen = 0
		for r := 3 * (box / 3); r < 3*(box/3)+3; r++ {
			for c := 4 * (box % 3); c < 4*(box%3)+4; c++ {
				if p[r][c] == 0 {
					continue
				}
				b = 1 << p[r][c]
				if seen&b == b {
					return false
				}
				seen |= b
			}
		}
	}

	return true
}
