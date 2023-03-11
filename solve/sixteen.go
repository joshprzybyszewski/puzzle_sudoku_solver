package solve

import (
	"errors"
	"time"

	"github.com/joshprzybyszewski/puzzle_sudoku_solver/model"
)

var (
	t0 time.Time
)

func Sixteen(
	puzzle model.Sixteen,
	timeout time.Duration,
) (model.Sixteen, error) {
	t0 = time.Now()
	s, ok := solveSixteenRow(&puzzle)
	if !ok {
		return model.Sixteen{}, errors.New(`did not solve`)
	}
	return s, nil
}

// func solveSixteen(
// 	s *model.Sixteen,
// ) (model.Sixteen, bool) {

// 	r, c := s.Best()
// 	if r > s.Size() {
// 		if s.IsSolved() {
// 			return *s, true
// 		}
// 		return model.Sixteen{}, false
// 	}
// 	for i := uint8(1); i <= s.Size(); i++ {
// 		cpy, ok := s.Place(r, c, i)
// 		if !ok {
// 			continue
// 		}
// 		solved, ok := solveSixteen(&cpy)
// 		if ok {
// 			return solved, true
// 		}
// 	}

// 	return model.Sixteen{}, false
// }

func solveSixteenRow(
	s *model.Sixteen,
) (model.Sixteen, bool) {

	if time.Since(t0) > 10*time.Second {
		return model.Sixteen{}, false
	}

	r := s.BestRow()
	if r > s.Size() {
		if s.IsSolved() {
			return *s, true
		}
		return model.Sixteen{}, false
	}

	rf := rowFilled{}
	rf.fillSixteenRow(s, r, 0, 0, func(*model.Sixteen) bool { return true })
	for i := 0; i < rf.lastIndex; i++ {
		base := *s
		if !rf.entries[i](&base) {
			continue
		}
		solved, ok := solveSixteenRow(&base)
		if ok {
			return solved, true
		}
	}

	return model.Sixteen{}, false
}

type rowFilled struct {
	/* 46656 = 6^6 */
	entries   [46656]func(*model.Sixteen) bool
	lastIndex int
}

func (rf *rowFilled) fillSixteenRow(
	s *model.Sixteen,
	r, c uint8,
	hasPlaced uint16,
	prev func(*model.Sixteen) bool,
) {

	for ; c < s.Size(); c++ {
		if s.IsSet(r, c) {
			continue
		}
		for i := uint8(1); i <= s.Size(); i++ {
			b := uint16(1) << (i - 1)
			if hasPlaced&b == b {
				continue
			}
			if !s.CanPlace(r, c, i) {
				continue
			}
			r := r
			c := c
			val := i
			my := func(s *model.Sixteen) bool {
				if !s.Place(r, c, val) {
					return false
				}
				return prev(s)
			}
			rf.fillSixteenRow(s, r, c+1, hasPlaced|b, my)
		}

		return
	}

	rf.entries[rf.lastIndex] = prev
	rf.lastIndex++
}
