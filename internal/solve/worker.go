package solve

import (
	"context"
)

type worker struct {
	state puzzle

	sendAnswer func(puzzle)
}

func newWorker(
	sendAnswer func(puzzle),
) worker {
	return worker{
		sendAnswer: sendAnswer,
	}
}

func (w *worker) process(
	ctx context.Context,
) {
	r := w.state.BestRow()
	if r > w.state.Size() {
		if w.state.IsSolved() {
			w.sendAnswer(w.state)
			return
		}
		return
	}

	f := newFiller()
	f.fillRow(
		&w.state,
		r,
		0, 0,
		pendingWrite{},
	)

	cpy := w.state
	for i := 0; i < f.lastIndex; i++ {
		if f.entries[i].applyRow(r, &w.state) {
			w.process(ctx)
		}
		w.state = cpy
	}
}
