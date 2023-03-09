package model

import "fmt"

type Iterator int

func (i Iterator) String() string {
	return fmt.Sprintf("%s %s", i.getSize(), i.GetDifficulty())
}

func (i Iterator) Valid() bool {
	return MinIterator <= i && i <= MaxIterator
}

func (i Iterator) getSize() string {
	if i < MinIterator || i > MaxIterator {
		return `invalid`
	}

	return `3x3`
}

func (i Iterator) GetDifficulty() Difficulty {
	if i < MinIterator || i > MaxIterator {
		return invalidDifficulty
	}

	if i >= Iterator(basic) && i <= Iterator(evil) {
		return Difficulty(i)
	}

	return invalidDifficulty
}
