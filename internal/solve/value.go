package solve

type value uint8

func (v value) bit() bits {
	return valsToBits[v]
}

var (
	valsToBits = [17]bits{
		0xFFFF, // ERROR
		1 << 0,
		1 << 1,
		1 << 2,
		1 << 3,
		1 << 4,
		1 << 5,
		1 << 6,
		1 << 7,
		1 << 8,
		1 << 9,
		1 << 10,
		1 << 11,
		1 << 12,
		1 << 13,
		1 << 14,
		1 << 15,
	}
)