package solve

type pendingWrite struct {
	indexes    [16]uint8
	vals       [16]value
	numIndexes uint8
}

func (pw pendingWrite) add(
	i uint8,
	v value,
) pendingWrite {
	pw.indexes[pw.numIndexes] = i
	pw.vals[pw.numIndexes] = v
	pw.numIndexes++
	return pw
}

func (pw *pendingWrite) applyRow(
	r uint8,
	s *puzzle,
) bool {

	for i := uint8(0); i < pw.numIndexes; i++ {
		if !s.Place(r, pw.indexes[i], pw.vals[i]) {
			return false
		}
	}
	return true
}

func (pw *pendingWrite) applyCol(
	c uint8,
	s *puzzle,
) bool {

	for i := uint8(0); i < pw.numIndexes; i++ {
		if !s.Place(pw.indexes[i], c, pw.vals[i]) {
			return false
		}
	}

	return true
}

type filler struct {
	/* 46656 = 6^6 */
	entries   [256]pendingWrite
	lastIndex int
}

func newFiller() filler {
	return filler{}
}

func (rf *filler) fillRow(
	s *puzzle,
	r, c uint8,
	hasPlaced bits,
	pw pendingWrite,
) {

	var val value
	var b bits

	for ; c < s.Size(); c++ {
		if s.IsSet(r, c) {
			hasPlaced |= s.Val(r, c).bit()
			continue
		}
		for val = 1; val <= value(s.Size()); val++ {
			b = val.bit()
			if hasPlaced&b == b {
				continue
			}
			if !s.CanPlace(r, c, val) {
				continue
			}

			rf.fillRow(
				s,
				r, c+1,
				hasPlaced|b,
				pw.add(c, val),
			)
		}

		return
	}

	rf.entries[rf.lastIndex] = pw
	rf.lastIndex++
}

func (rf *filler) fillCol(
	s *puzzle,
	r, c uint8,
	hasPlaced bits,
	pw pendingWrite,
) {

	var val value
	var b bits

	for ; r < s.Size(); r++ {
		if s.IsSet(r, c) {
			hasPlaced |= s.Val(r, c).bit()
			continue
		}
		for val = 1; val <= value(s.Size()); val++ {
			b = val.bit()
			if hasPlaced&b == b {
				continue
			}
			if !s.CanPlace(r, c, val) {
				continue
			}
			rf.fillCol(
				s,
				r+1, c,
				hasPlaced|b,
				pw.add(r, val),
			)
		}

		return
	}

	rf.entries[rf.lastIndex] = pw
	rf.lastIndex++
}
