package solve

import (
	"context"
	"fmt"
	"runtime"
	"runtime/debug"
	"strings"
	"time"
)

const (
	// maxNumWorkers = 1
	maxNumWorkers = 8
)

type workforce struct {
	solution chan puzzle

	work chan puzzle

	workers [maxNumWorkers]worker
}

func newWorkforce() workforce {
	wf := workforce{
		solution: make(chan puzzle, 1),
		work:     make(chan puzzle, runtime.NumCPU()),
	}

	for i := range wf.workers {
		wf.workers[i] = newWorker(
			func(sol puzzle) {
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
	puzz puzzle,
) (puzzle, error) {

	if puzz.IsSolved() {
		return puzz, nil
	}

	go w.sendWork(ctx, puzz)

	select {
	case <-ctx.Done():
		return puzzle{}, fmt.Errorf("Ran out of time.")
	case sol, ok := <-w.solution:
		if !ok {
			return puzzle{}, fmt.Errorf("did not find the solution")
		}
		return sol, nil
	}
}

func (w *workforce) sendWork(
	ctx context.Context,
	initial puzzle,
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

	w.work <- initial
	if len(w.workers) == 1 {
		// if there is only one worker, then we _need_ the initial state to be solved.
		select {
		case <-ctx.Done():
			return
		}
	}

	c := initial.BestCol()
	if c > initial.Size() {
		if initial.IsSolved() {
			w.solution <- initial
		}
		return
	}

	f := newFiller()
	f.fillCol(
		&initial,
		0, c,
		0,
		pendingWrite{},
	)

	cpy := initial
	for i := 0; i < f.lastIndex; i++ {
		if f.entries[i].applyCol(c, &cpy) {
			w.work <- cpy
		}
		cpy = initial
	}
}
