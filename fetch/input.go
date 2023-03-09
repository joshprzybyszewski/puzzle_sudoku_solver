package fetch

import (
	"fmt"

	"github.com/joshprzybyszewski/slitherlink/model"
)

type input struct {
	ID    string
	param string
	task  string

	iter model.Iterator
}

func (i input) String() string {
	return fmt.Sprintf("Puzzle %s (Iter: %d, Difficulty: %s)",
		i.ID,
		i.iter,
		i.iter.GetDifficulty(),
	)
}

func (i input) Task() string {
	return i.task
}

func (i input) ToClassic() model.Classic {
	var r, c int

	var output model.Classic

	for _, b := range i.task {
		if b >= '0' && b <= '9' {
			output[r][c] = uint8(b - '0')
		} else if b != '_' {
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
