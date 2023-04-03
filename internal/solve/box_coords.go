package solve

func init() {
	for r := uint8(0); r < 9; r++ {
		for c := uint8(0); c < 9; c++ {
			defaultNineBoxLookups[r][c] = (3 * (r / 3)) + (c / 3)
		}
	}
	for i := range defaultNineBoxes {
		defaultNineBoxes[i] = createBox(uint8(i), defaultNineBoxLookups, 9)
	}

	for r := uint8(0); r < 12; r++ {
		for c := uint8(0); c < 12; c++ {
			defaultTwelveBoxLookups[r][c] = (3 * (r / 3)) + (c / 4)
		}
	}
	for i := range defaultTwelveBoxes {
		defaultTwelveBoxes[i] = createBox(uint8(i), defaultTwelveBoxLookups, 12)
	}

	for r := uint8(0); r < 16; r++ {
		for c := uint8(0); c < 16; c++ {
			defaultSixteenBoxLookups[r][c] = (4 * (r / 4)) + (c / 4)
		}
	}
	for i := range defaultTwelveBoxes {
		defaultSixteenBoxes[i] = createBox(uint8(i), defaultSixteenBoxLookups, 16)
	}
}

var (
	defaultNineBoxLookups    [16][16]uint8
	defaultTwelveBoxLookups  [16][16]uint8
	defaultSixteenBoxLookups [16][16]uint8

	defaultNineBoxes    [16]box
	defaultTwelveBoxes  [16]box
	defaultSixteenBoxes [16]box
)

func createBox(
	bi uint8,
	indexes [16][16]uint8,
	size uint8,
) box {
	var b box
	var num, c uint8
	for r := uint8(0); r < size; r++ {
		for c = uint8(0); c < size; c++ {
			if indexes[r][c] != bi {
				continue
			}
			b[num] = boxCell{
				row: r,
				col: c,
			}
			num++
		}
	}
	return b
}

type box [16]boxCell

type boxCell struct {
	row uint8
	col uint8
}
