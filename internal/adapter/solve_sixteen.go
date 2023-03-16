package adapter

import (
	"context"
	"time"

	smodel "github.com/joshprzybyszewski/puzzle_sudoku_solver/internal/model"
	"github.com/joshprzybyszewski/puzzle_sudoku_solver/internal/solve"
	"github.com/joshprzybyszewski/puzzler/model"
)

func solveSixteen(
	g *model.Game,
	timeout time.Duration,
) error {
	sud := smodel.NewSixteen(g.Task)

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
