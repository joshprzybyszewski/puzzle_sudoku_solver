package model

import "fmt"

type Iterator int

func (i Iterator) String() string {
	r, c := i.GetSize()
	if i <= MaxIteratorStandard {
		return fmt.Sprintf("%dx%d %s", r, c, i.GetDifficulty())
	}
	return fmt.Sprintf("jigsaw %dx%d %s", r, c, i.GetDifficulty())
}

func (i Iterator) GetSize() (uint8, uint8) {
	if i < MinIteratorStandard || i > MaxIteratorJigsaw {
		return 0, 0
	}
	if i > MaxIteratorStandard && i < MinIteratorJigsaw {
		// unexpected
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

	switch i {
	case 100, 101, 102:
		return 5, 5
	case 103, 104, 105:
		return 7, 7
	case 106, 107, 108:
		return 9, 9
	}

	return 0, 0
}

func (i Iterator) GetDifficulty() Difficulty {
	if i < MinIteratorStandard || i > MaxIteratorJigsaw {
		return invalidDifficulty
	}
	if i > MaxIteratorStandard && i < MinIteratorJigsaw {
		// unexpected
		return invalidDifficulty
	}

	if i <= MaxIteratorStandard {
		if i >= Iterator(basic) && i <= Iterator(evil) {
			return Difficulty(i)
		}

		return advanced
	}

	switch i {
	case 100, 103, 106:
		return easy
	case 101, 104, 107:
		return intermediate
	case 102, 105, 108:
		return advanced
	}
	return invalidDifficulty
}
