package model

import "strconv"

type Difficulty uint8

const (
	invalidDifficulty Difficulty = 42

	basic        Difficulty = 0
	easy         Difficulty = 1
	intermediate Difficulty = 2
	advanced     Difficulty = 3
	extreme      Difficulty = 4
	evil         Difficulty = 5
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
