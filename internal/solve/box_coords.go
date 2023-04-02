package solve

func init() {
	num := [16]int{}

	for r := uint8(0); r < 9; r++ {
		for c := uint8(0); c < 9; c++ {
			bci := (3 * (r / 3)) + (c / 3)
			defaultNineBoxLookups[r][c] = bci
			defaultNineBoxes[bci][num[bci]] = boxCell{
				row: r,
				col: c,
			}
			num[bci]++
		}
	}

	num = [16]int{}
	for r := uint8(0); r < 12; r++ {
		for c := uint8(0); c < 12; c++ {
			bci := (3 * (r / 3)) + (c / 4)
			defaultTwelveBoxLookups[r][c] = bci
			defaultTwelveBoxes[bci][num[bci]] = boxCell{
				row: r,
				col: c,
			}
			num[bci]++
		}
	}

	num = [16]int{}
	for r := uint8(0); r < 16; r++ {
		for c := uint8(0); c < 16; c++ {
			bci := (4 * (r / 4)) + (c / 4)
			defaultSixteenBoxLookups[r][c] = bci
			defaultSixteenBoxes[bci][num[bci]] = boxCell{
				row: r,
				col: c,
			}
			num[bci]++
		}
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

type box [16]boxCell

type boxCell struct {
	row uint8
	col uint8
}
