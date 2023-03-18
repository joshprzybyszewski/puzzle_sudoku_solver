package solve

import (
	"context"
	"fmt"
	"runtime"
	"runtime/debug"
	"strings"
	"time"

	"github.com/joshprzybyszewski/puzzle_sudoku_solver/internal/model"
)

const (
	// maxNumWorkers = 1
	maxNumWorkers = 8
)

type worker struct {
	state model.Sixteen

	sendAnswer func(model.Sixteen)
}

func newWorker(
	sendAnswer func(model.Sixteen),
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
			fmt.Printf("solved\n%s\n", &w.state)
			w.sendAnswer(w.state)
			return
		}
		fmt.Printf("not solved\n%s\n", &w.state)
		return
	}
	// fmt.Printf("processing\n%s\n", &w.state)

	rf := rowFilled{}
	rf.fillSixteenRow(
		&w.state,
		r,
		0, 0,
		func(*model.Sixteen) bool { return true },
	)

	cpy := w.state
	for i := 0; i < rf.lastIndex; i++ {
		if rf.entries[i](&w.state) {
			w.process(ctx)
		}
		w.state = cpy
	}
}

type workforce struct {
	solution chan model.Sixteen

	work chan model.Sixteen

	workers [maxNumWorkers]worker
}

func newWorkforce() workforce {
	wf := workforce{
		solution: make(chan model.Sixteen, 1),
		work:     make(chan model.Sixteen, runtime.NumCPU()),
	}

	for i := range wf.workers {
		wf.workers[i] = newWorker(
			func(sol model.Sixteen) {
				defer func() {
					// if the solution channel has been closed, then don't do anything.
					_ = recover()
				}()
				wf.solution <- sol
			},
		)
	}

	return wf
}

func (w *workforce) start(
	ctx context.Context,
) {
	max := runtime.NumCPU()

	for i := range w.workers {
		i := i
		if i >= max {
			break
		}
		go w.startWorker(
			ctx,
			&w.workers[i],
			i,
		)
	}
}

func (w *workforce) startWorker(
	ctx context.Context,
	worker *worker,
	id int,
) {
	var ok bool

	idleLogDur := 500 * time.Millisecond

	for {
		select {
		case <-ctx.Done():
			return
		case <-time.After(idleLogDur):
			fmt.Printf("Worker %d is idle...\n", id)
			idleLogDur += idleLogDur
		case worker.state, ok = <-w.work:
			if !ok {
				return
			}
			worker.process(ctx)
			idleLogDur = 500 * time.Millisecond
		}
	}
}

func (w *workforce) stop() {
	close(w.work)
	close(w.solution)
}

func (w *workforce) solve(
	ctx context.Context,
	puzz model.Sixteen,
) (model.Sixteen, error) {

	if puzz.IsSolved() {
		return puzz, nil
	}

	go w.sendWork(ctx, puzz)

	select {
	case <-ctx.Done():
		return model.Sixteen{}, fmt.Errorf("Ran out of time.")
	case sol, ok := <-w.solution:
		if !ok {
			return model.Sixteen{}, fmt.Errorf("did not find the solution")
		}
		return sol, nil
	}
}

func (w *workforce) sendWork(
	ctx context.Context,
	initial model.Sixteen,
) {
	defer func() {
		// if the work channel has been closed, then don't do anything.
		r := recover()
		if r != nil {
			if strings.Contains(fmt.Sprintf("%+v", r), "send on closed channel") {
				return
			}
			fmt.Printf("caught: %+v\n", r)
			fmt.Printf("%s\n", debug.Stack())
		}
	}()

	if ctx.Err() != nil {
		return
	}

	// TODO send the initial state, then find a way to send similar, but different work.
	// w.work <- initial
	// if len(w.workers) == 1 {
	// 	// if there is only one worker, then we _need_ the initial state to be solved.
	// 	select {
	// 	case <-ctx.Done():
	// 		return
	// 	}
	// }

	r := initial.BestRow()
	if r > initial.Size() {
		if initial.IsSolved() {
			w.solution <- initial
			return
		}
		panic(`what happened`)
	}

	// TODO find a way to fill "the second best row" or something.
	rf := rowFilled{}
	rf.fillSixteenRow(
		&initial,
		r, 0,
		0,
		func(*model.Sixteen) bool { return true },
	)

	fmt.Printf("inital state has %d permutations\n", rf.lastIndex)
	cpy := initial
	for i := 0; i < rf.lastIndex; i++ {
		if rf.entries[i](&cpy) {
			w.work <- cpy
		}
		cpy = initial
	}
}