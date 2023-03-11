package adapter

import (
	"time"

	smodel "github.com/joshprzybyszewski/puzzle_sudoku_solver/model"
	"github.com/joshprzybyszewski/puzzle_sudoku_solver/solve"
	"github.com/joshprzybyszewski/puzzler/model"
)

func solveSixteen(
	g *model.Game,
	timeout time.Duration,
) error {
	iter := smodel.Iterator(g.Iterator)
	sud := convertSixteenTask(iter, g.Task)

	sol, err := solve.Sixteen(
		sud,
		timeout,
	)
	if err != nil {
		return err
	}

	g.Answer = convertSixteenAnswer(sol)
	return nil
}

func convertSixteenTask(
	iter smodel.Iterator,
	task model.Task,
) smodel.Sixteen {
	var r, c int

	var output smodel.Sixteen

	var b byte
	for i := 0; i < len(task); i++ {
		b = task[i]
		if b > '0' && b <= '9' {
			output[r][c] = uint8(b - '0')
			if output[r][c] == 1 && task[i+1] >= '0' && task[i+1] <= '6' {
				output[r][c] = 10 + uint8(task[i+1]-'0')
				i++
			}
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

func convertSixteenAnswer(
	p smodel.Sixteen,
) model.Answer {
	output := make([]byte, 0, len(p)*len(p[0])*2)

	for r := range p {
		for c := range p[r] {
			if p[r][c] >= 10 {
				output = append(output, '1')
				p[r][c] -= 10
			}
			output = append(output, '0'+p[r][c])
			output = append(output, ',')
		}
	}

	// omit the last comma
	return model.Answer(output[:len(output)-1])
}
