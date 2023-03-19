package solve

import (
	"context"

	"github.com/joshprzybyszewski/puzzle_sudoku_solver/internal/model"
)

func Sixteen(
	ctx context.Context,
	puzzle model.Sixteen,
) (model.Sixteen, error) {

	wf := newWorkforce()
	wf.start(ctx)
	defer wf.stop()

	return wf.solve(ctx, puzzle)
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

	var i uint8
	var b uint16

	for ; c < s.Size(); c++ {
		if s.IsSet(r, c) {
			hasPlaced |= valsToBits[s.Val(r, c)]
			continue
		}
		for i = uint8(1); i <= s.Size(); i++ {
			b = valsToBits[i]
			if hasPlaced&b == b {
				continue
			}
			if !s.CanPlace(r, c, i) {
				continue
			}
			r := r
			c := c
			val := i
			rf.fillSixteenRow(
				s,
				r, c+1,
				hasPlaced|b,
				func(sub *model.Sixteen) bool {
					if !sub.Place(r, c, val) {
						return false
					}
					return prev(sub)
				},
			)
		}

		return
	}

	rf.entries[rf.lastIndex] = prev
	rf.lastIndex++
}

func (rf *rowFilled) fillSixteenCol(
	s *model.Sixteen,
	r, c uint8,
	hasPlaced uint16,
	prev func(*model.Sixteen) bool,
) {

	var i uint8
	var b uint16

	for ; r < s.Size(); r++ {
		if s.IsSet(r, c) {
			hasPlaced |= valsToBits[s.Val(r, c)]
			continue
		}
		for i = uint8(1); i <= s.Size(); i++ {
			b = valsToBits[i]
			if hasPlaced&b == b {
				continue
			}
			if !s.CanPlace(r, c, i) {
				continue
			}
			r := r
			c := c
			val := i
			rf.fillSixteenCol(
				s,
				r+1, c,
				hasPlaced|b,
				func(sub *model.Sixteen) bool {
					if !sub.Place(r, c, val) {
						return false
					}
					return prev(sub)
				},
			)
		}

		return
	}

	rf.entries[rf.lastIndex] = prev
	rf.lastIndex++
}
