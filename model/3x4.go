package model

type Twelve [12][12]uint8

func (p Twelve) String() string {
	output := make([]byte, 0, len(p)*len(p[0])*2)

	for r := range p {
		for c := range p[r] {
			if p[r][c] > 9 {
				output = append(output, 'A'+(p[r][c]-9))
			} else {
				output = append(output, '0'+p[r][c])
			}
			output = append(output, ',')
		}
		output = append(output, '\n')
	}

	// omit the last comma
	return string(output[:len(output)-1])

}
