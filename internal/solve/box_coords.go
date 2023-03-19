package solve

func init() {
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
	// TODO add other sizes
	sixteenBoxCoords [16][16]boxCoords
)

type boxCoords struct {
	startR, stopR uint8
	startC, stopC uint8
}
