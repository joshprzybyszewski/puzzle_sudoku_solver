package model

import "strconv"

type Difficulty uint8

const (
	invalidDifficulty Difficulty = 0

	basic        Difficulty = 1
	easy         Difficulty = 2
	intermediate Difficulty = 3
	advanced     Difficulty = 4
	extreme      Difficulty = 5
	evil         Difficulty = 6
)

func (d Difficulty) String() string {
	switch d {
	case basic:
		return `basic`
	case easy:
		return `easy`
	case intermediate:
		return `intermediate`
	case advanced:
		return `advanced`
	case extreme:
		return `extreme`
	case evil:
		return `evil`
	default:
		return strconv.Itoa(int(d))
	}
}
