package solve

import (
	"context"
	"errors"
	"runtime"

	"github.com/joshprzybyszewski/puzzle_sudoku_solver/model"
)

func Sixteen(
	ctx context.Context,
	puzzle model.Sixteen,
) (model.Sixteen, error) {
	r := puzzle.BestRow()
	if r > puzzle.Size() {
		if puzzle.IsSolved() {
			return puzzle, nil
		}
		return model.Sixteen{}, errors.New(`bad initial state`)
	}

	rf := rowFilled{}
	rf.fillSixteenRow(&puzzle, r, 0, 0, func(i model.Sixteen) (model.Sixteen, bool) { return i, true })

	work := make(chan func(model.Sixteen) (model.Sixteen, bool), rf.lastIndex)
	solution := make(chan model.Sixteen)
	defer close(solution)
	defer close(work)
	start := puzzle
	for i := 0; i < runtime.NumCPU(); i++ {
		go func() {
			for {
				select {
				case <-ctx.Done():
					return
				case fn, ok := <-work:
					if !ok {
						return
					}
					p, ok := fn(start)
					if !ok {
						continue
					}
					solved, ok := solveSixteenRow(ctx, p)
					if ok {
						solution <- solved
					}
				}
			}
		}()
	}

	for i := 0; i < rf.lastIndex; i++ {
		work <- rf.entries[i]
	}

	select {
	case <-ctx.Done():
		return model.Sixteen{}, ctx.Err()
	case s, ok := <-solution:
		if !ok {
			return model.Sixteen{}, errors.New(`solution not found`)
		}
		return s, nil
	}
}

func solveSixteenRow(
	ctx context.Context,
	s model.Sixteen,
) (model.Sixteen, bool) {

	r := s.BestRow()
	if r > s.Size() {
		if s.IsSolved() {
			return s, true
		}
		return model.Sixteen{}, false
	}

	rf := rowFilled{}
	rf.fillSixteenRow(&s, r, 0, 0, func(i model.Sixteen) (model.Sixteen, bool) { return i, true })
	if ctx.Err() != nil {
		return model.Sixteen{}, false
	}

	start := s
	for i := 0; i < rf.lastIndex; i++ {
		e, ok := rf.entries[i](start)
		if !ok {
			continue
		}
		solved, ok := solveSixteenRow(ctx, e)
		if ok {
			return solved, true
		}
	}

	return model.Sixteen{}, false
}

type rowFilled struct {
	/* 46656 = 6^6 */
	entries   [46656]func(model.Sixteen) (model.Sixteen, bool)
	lastIndex int
}

func (rf *rowFilled) fillSixteenRow(
	s *model.Sixteen,
	r, c uint8,
	hasPlaced uint16,
	prev func(model.Sixteen) (model.Sixteen, bool),
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
			my := func(s model.Sixteen) (model.Sixteen, bool) {
				if !s.Place(r, c, val) {
					return model.Sixteen{}, false
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
