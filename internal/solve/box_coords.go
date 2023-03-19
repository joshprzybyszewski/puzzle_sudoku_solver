package solve

func init() {
	for r := range nineBoxCoords {
		for c := range nineBoxCoords {
			nineBoxCoords[r][c].startR = 3 * (uint8(r) / 3)
			nineBoxCoords[r][c].stopR = nineBoxCoords[r][c].startR + 3

			nineBoxCoords[r][c].startC = 3 * (uint8(c) / 3)
			nineBoxCoords[r][c].stopC = nineBoxCoords[r][c].startC + 3
		}
	}

	for r := range twelveBoxCoords {
		for c := range twelveBoxCoords {
			twelveBoxCoords[r][c].startR = 3 * (uint8(r) / 3)
			twelveBoxCoords[r][c].stopR = twelveBoxCoords[r][c].startR + 3

			twelveBoxCoords[r][c].startC = 4 * (uint8(c) / 4)
			twelveBoxCoords[r][c].stopC = twelveBoxCoords[r][c].startC + 4
		}
	}

	for r := range sixteenBoxCoords {
		for c := range sixteenBoxCoords {
			sixteenBoxCoords[r][c].startR = 4 * (uint8(r) / 4)
			sixteenBoxCoords[r][c].stopR = sixteenBoxCoords[r][c].startR + 4

			sixteenBoxCoords[r][c].startC = 4 * (uint8(c) / 4)
			sixteenBoxCoords[r][c].stopC = sixteenBoxCoords[r][c].startC + 4
		}
	}
}

var (
	nineBoxCoords    [9][9]boxCoords
	twelveBoxCoords  [12][12]boxCoords
	sixteenBoxCoords [16][16]boxCoords

	invalidBoxCoords = boxCoords{
		startR: 17,
		stopR:  18,
		startC: 19,
		stopC:  20,
	}
)

type boxCoords struct {
	startR, stopR uint8
	startC, stopC uint8
}
