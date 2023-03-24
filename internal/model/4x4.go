package model

import (
	pmodel "github.com/joshprzybyszewski/puzzler/model"
)

type Sixteen [16][16]uint8

func NewSixteen(
	task pmodel.Task,
) Sixteen {
	puzz := Sixteen{}

	var r, c int

	var b byte
	for i := 0; i < len(task); i++ {
		b = task[i]
		if b == '_' {
			continue
		}

		if b >= '0' && b <= '9' {
			if b == '1' && i+1 < len(task) && task[i+1] >= '0' && task[i+1] <= '6' {
				puzz[r][c] = 10 + task[i+1] - '0'
				i++
			} else {
				puzz[r][c] = b - '0'
			}
		} else {
			c += int(b - 'a')
		}

		c++

		if c >= len(puzz) {
			r += (c / len(puzz))
			c %= len(puzz)
		}
	}

	return puzz
}

func (p Sixteen) String() string {
	output := make([]byte, 0, len(p)*len(p[0])*2)

	for r := range p {
		for c := range p[r] {
			output = append(output, '0'+p[r][c])
			output = append(output, ',')
		}
		output = append(output, '\n')
	}

	// omit the last comma
	return string(output[:len(output)-1])

}
