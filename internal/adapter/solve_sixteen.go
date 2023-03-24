package adapter

import (
	"context"

	smodel "github.com/joshprzybyszewski/puzzle_sudoku_solver/internal/model"
	"github.com/joshprzybyszewski/puzzle_sudoku_solver/internal/solve"
	"github.com/joshprzybyszewski/puzzler/model"
)

func solveSixteen(
	ctx context.Context,
	g *model.Game,
) error {
	sud := smodel.NewSixteen(g.Task)

	sol, err := solve.Sixteen(
		ctx,
		sud,
	)
	if err != nil {
		return err
	}

	g.Answer = convertSixteenAnswer(&sol)
	return nil
}

func convertSixteenAnswer(
	p *smodel.Sixteen,
) model.Answer {
	output := make([]byte, 0, len(p)*len(p)*2)

	var c int
	var v uint8

	for r := range p {
		for c = range p[r] {
			v = p[r][c]
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
