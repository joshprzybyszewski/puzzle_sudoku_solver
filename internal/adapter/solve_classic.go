package adapter

import (
	"time"

	smodel "github.com/joshprzybyszewski/puzzle_sudoku_solver/internal/model"
	"github.com/joshprzybyszewski/puzzle_sudoku_solver/internal/solve"
	"github.com/joshprzybyszewski/puzzler/model"
)

func solveClassic(
	g *model.Game,
	timeout time.Duration,
) error {
	iter := smodel.Iterator(g.Iterator)
	sud := convertClassicTask(iter, g.Task)

	sol, err := solve.Classic(
		sud,
		timeout,
	)
	if err != nil {
		return err
	}

	g.Answer = convertClassicAnswer(sol)
	return nil
}

func convertClassicTask(
	iter smodel.Iterator,
	task model.Task,
) smodel.Classic {
	var r, c int

	var output smodel.Classic

	for _, b := range task {
		if b > '0' && b <= '9' {
			output[r][c] = uint8(b - '0')
		} else if b == '_' {
			continue
		} else {
			c += int(b - 'a')
		}

		c++

		if c >= len(output[r]) {
			r += (c / len(output[r]))
			c %= len(output[r])
		}
	}

	return output
}

func convertClassicAnswer(
	p smodel.Classic,
) model.Answer {
	output := make([]byte, 0, len(p)*len(p[0])*2)

	for r := range p {
		for c := range p[r] {
			output = append(output, '0'+p[r][c])
			output = append(output, ',')
		}
	}

	// omit the last comma
	return model.Answer(output[:len(output)-1])
}
