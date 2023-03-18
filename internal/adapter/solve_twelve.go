package adapter

import (
	"context"

	smodel "github.com/joshprzybyszewski/puzzle_sudoku_solver/internal/model"
	"github.com/joshprzybyszewski/puzzle_sudoku_solver/internal/solve"
	"github.com/joshprzybyszewski/puzzler/model"
)

func solveTwelve(
	ctx context.Context,
	g *model.Game,
) error {
	sud := ConvertTwelveTask(g.Task)

	sol, err := solve.Twelve(
		ctx,
		sud,
	)
	if err != nil {
		return err
	}

	g.Answer = convertTwelveAnswer(sol)
	return nil
}

func convertTwelveAnswer(
	p smodel.Twelve,
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
