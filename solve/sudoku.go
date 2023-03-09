package solve

import (
	"errors"

	"github.com/joshprzybyszewski/slitherlink/model"
)

func SolveClassic(
	puzzle model.Classic,
) (model.Classic, error) {
	s, ok := solve(puzzle)
	if !ok {
		return model.Classic{}, errors.New(`did not solve`)
	}
	return s, nil
}

func solve(
	s model.Classic,
) (model.Classic, bool) {
	for r := range s {
		for c := range s[r] {
			if s[r][c] != 0 {
				continue
			}

			for i := 1; i <= len(s); i++ {
				s[r][c] = uint8(i)
				solved, ok := solve(s)
				if ok {
					return solved, true
				}
			}

			return model.Classic{}, false
		}
	}

	if !isValid(s) {
		return model.Classic{}, false
	}

	return s, true
}

func isValid(
	p model.Classic,
) bool {
	var seen, b uint16
	// check each row that it has all the numbers
	for r := 0; r < len(p); r++ {
		seen = 0
		for c := range p[r] {
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
			b = 1 << p[r][c]
			if seen&b == b {
				return false
			}
			seen |= b
		}
	}

	// TODO
	// check each box that it has all the numbers
	// for c := 0; c < len(p[0]); c++ {
	// 	seen = 0
	// 	for r := 0; r < len(p); r++ {
	// 		b = 1 << p[r][c]
	// 		if seen&b == b {
	// 			return false
	// 		}
	// 		seen |= b
	// 	}
	// }

	return true
}
