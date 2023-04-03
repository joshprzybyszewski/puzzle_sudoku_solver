package adapter

import (
	"context"

	smodel "github.com/joshprzybyszewski/puzzle_sudoku_solver/internal/model"
	"github.com/joshprzybyszewski/puzzle_sudoku_solver/internal/solve"
	"github.com/joshprzybyszewski/puzzler/model"
)

func solveJigsaw(
	ctx context.Context,
	g *model.Game,
) error {
	input, boxNums := convertJigsawTask(
		g.Task,
		smodel.Iterator(g.Iterator),
	)

	sol, err := solve.Jigsaw(
		ctx,
		input,
		boxNums,
	)
	if err != nil {
		return err
	}

	g.Answer = convertJigsawAnswer(sol)
	return nil
}

func convertJigsawTask(
	task model.Task,
	iterator smodel.Iterator,
) ([][]uint8, [][]uint8) {
	var r, c int

	numRows, numCols := iterator.GetSize()
	nci := int(numCols)
	output := make([][]uint8, numRows)
	output[0] = make([]uint8, numCols)

	var boxNums [][]uint8

	for i, b := range task {
		if b == ';' {
			boxNums = convertJigsawBoxNums(string(task[i+1:]), iterator)
			break
		}

		if b > '0' && b <= '9' {
			output[r][c] = uint8(b - '0')
		} else if b == '_' {
			continue
		} else {
			c += int(b - 'a')
		}

		c++

		if c >= nci {
			r += (c / nci)
			c %= nci
			if r < len(output) {
				output[r] = make([]uint8, numCols)
			}
		}
	}

	return output, boxNums
}

func convertJigsawBoxNums(
	taskSuffix string,
	iterator smodel.Iterator,
) [][]uint8 {
	var r, c int

	numRows, numCols := iterator.GetSize()
	nci := int(numCols)
	boxNums := make([][]uint8, numRows)
	boxNums[0] = make([]uint8, numCols)

	for _, b := range taskSuffix {
		if b == ',' {
			continue
		}

		if b > '0' && b <= '9' {
			boxNums[r][c] = uint8(b - '0')
		}

		c++

		if c >= nci {
			r += (c / nci)
			c %= nci
			if r < len(boxNums) {
				boxNums[r] = make([]uint8, numCols)
			}
		}
	}

	return boxNums
}

func convertJigsawAnswer(
	p [][]uint8,
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
