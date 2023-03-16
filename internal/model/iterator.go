package model

import "fmt"

type Iterator int

func (i Iterator) String() string {
	r, c := i.GetSize()
	return fmt.Sprintf("%dx%d %s", r, c, i.GetDifficulty())
}

func (i Iterator) GetSize() (uint8, uint8) {
	if i < MinIterator || i > MaxIterator {
		return 0, 0
	}

	if i >= Iterator(basic) && i <= Iterator(evil) {
		return 3, 3
	}

	if i == 6 {
		return 3, 4
	}

	if i == 7 {
		return 4, 4
	}

	return 0, 0
}

func (i Iterator) GetDifficulty() Difficulty {
	if i < MinIterator || i > MaxIterator {
		return invalidDifficulty
	}

	if i >= Iterator(basic) && i <= Iterator(evil) {
		return Difficulty(i)
	}

	return advanced
}
