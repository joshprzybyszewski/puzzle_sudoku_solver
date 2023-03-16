package adapter

import (
	smodel "github.com/joshprzybyszewski/puzzle_sudoku_solver/internal/model"
	"github.com/joshprzybyszewski/puzzler/model"
)

func ConvertClassicTask(
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
