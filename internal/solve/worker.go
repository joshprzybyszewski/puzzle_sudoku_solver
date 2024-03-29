package solve

type worker struct {
	sendAnswer func(puzzle)

	state puzzle
}

func newWorker(
	sendAnswer func(puzzle),
) worker {
	return worker{
		sendAnswer: sendAnswer,
	}
}

func (w *worker) process() {
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
	for i := 0; i < f.lastIndex && i < len(f.entries); i++ {
		if f.entries[i].applyRow(r, &w.state) {
			w.process()
		}
		w.state = cpy
	}
	for _, pw := range f.extras {
		if pw.applyRow(r, &w.state) {
			w.process()
		}
		w.state = cpy
	}
}
