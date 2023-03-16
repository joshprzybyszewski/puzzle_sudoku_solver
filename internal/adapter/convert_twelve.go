package adapter

import (
	smodel "github.com/joshprzybyszewski/puzzle_sudoku_solver/internal/model"
	"github.com/joshprzybyszewski/puzzler/model"
)

func ConvertTwelveTask(
	task model.Task,
) smodel.Twelve {
	var r, c int

	var output smodel.Twelve

	var b byte
	for i := 0; i < len(task); i++ {
		b = task[i]
		if b > '0' && b <= '9' {
			output[r][c] = uint8(b - '0')
			if output[r][c] == 1 && task[i+1] >= '0' && task[i+1] <= '2' {
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
