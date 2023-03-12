package adapter

import (
	"context"
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

	ctx, cancelFn := context.WithTimeout(context.Background(), timeout)
	defer cancelFn()

	sol, err := solve.Sixteen(
		ctx,
		sud,
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
	var r, c uint8

	output := smodel.NewSixteen()

	var b byte
	for i := 0; i < len(task); i++ {
		b = task[i]
		if b == '_' {
			continue
		}

		if b >= '0' && b <= '9' {
			if b == '1' && i+1 < len(task) && task[i+1] >= '0' && task[i+1] <= '6' {
				output.InitialPlace(r, c, 10+uint8(task[i+1]-'0'))
				i++
			} else {
				output.InitialPlace(r, c, uint8(b-'0'))
			}
		} else {
			c += uint8(b - 'a')
		}

		c++

		if c >= output.Size() {
			r += (c / output.Size())
			c %= output.Size()
		}
	}

	return output
}

func convertSixteenAnswer(
	p smodel.Sixteen,
) model.Answer {
	output := make([]byte, 0, int(p.Size())*int(p.Size())*2)

	for r := uint8(0); r < p.Size(); r++ {
		for c := uint8(0); c < p.Size(); c++ {
			v := p.Val(r, c)
			if v >= 10 {
				output = append(output, '1')
				v -= 10
			}
			output = append(output, '0'+v)
			output = append(output, ',')
		}
	}

	// omit the last comma
	return model.Answer(output[:len(output)-1])
}
