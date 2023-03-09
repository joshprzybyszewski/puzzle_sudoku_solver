package model

type Classic [9][9]uint8

func (p Classic) ToAnswer() string {
	output := make([]byte, 0, len(p)*len(p[0])*2)

	for r := range p {
		for c := range p[r] {
			output = append(output, p[r][c], ',')
		}
	}

	// omit the last comma
	return string(output[:len(output)-1])

}
