package solve

type pendingWrite struct {
	r, c uint8
	val  value

	prev *pendingWrite
}

func (pw *pendingWrite) apply(s *puzzle) bool {
	if pw.val == 0 {
		return true
	}
	if !s.Place(pw.r, pw.c, pw.val) {
		return false
	}
	return pw.prev.apply(s)
}

type filler struct {
	/* 46656 = 6^6 */
	entries   [46656]pendingWrite
	lastIndex int
}

func newFiller() filler {
	return filler{}
}

func (rf *filler) fillRow(
	s *puzzle,
	r, c uint8,
	hasPlaced bits,
	prev pendingWrite,
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
				pendingWrite{
					r:    r,
					c:    c,
					val:  val,
					prev: &prev,
				},
			)
		}

		return
	}

	rf.entries[rf.lastIndex] = prev
	rf.lastIndex++
}

func (rf *filler) fillCol(
	s *puzzle,
	r, c uint8,
	hasPlaced bits,
	prev pendingWrite,
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
				pendingWrite{
					r:    r,
					c:    c,
					val:  val,
					prev: &prev,
				},
			)
		}

		return
	}

	rf.entries[rf.lastIndex] = prev
	rf.lastIndex++
}
