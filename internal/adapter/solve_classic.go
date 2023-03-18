package adapter

import (
	"context"

	smodel "github.com/joshprzybyszewski/puzzle_sudoku_solver/internal/model"
	"github.com/joshprzybyszewski/puzzle_sudoku_solver/internal/solve"
	"github.com/joshprzybyszewski/puzzler/model"
)

func solveClassic(
	ctx context.Context,
	g *model.Game,
) error {
	sud := ConvertClassicTask(g.Task)

	sol, err := solve.Classic(
		ctx,
		sud,
	)
	if err != nil {
		return err
	}

	g.Answer = convertClassicAnswer(sol)
	return nil
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
