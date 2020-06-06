package main

const (
	numStart       = 0x0030
	numEnd         = 0x0039
	alphabetLStart = 0x0041
	alphabetLEnd   = 0x005A
	alphabetSStart = 0x0061
	alphabetSEnd   = 0x007A
)

func shift(chars string, num int, reverse bool) string {
	rc := []rune(chars)
	size := len(rc)
	shifted := make([]rune, size)

	for i, c := range rc {
		var s rune
		if c >= numStart && c <= numEnd {
			s = shiftCore(c, numStart, numEnd, num, reverse)
		} else if c >= alphabetLStart && c <= alphabetLEnd {
			s = shiftCore(c, alphabetLStart, alphabetLEnd, num, reverse)
		} else if c >= alphabetSStart && c <= alphabetSEnd {
			s = shiftCore(c, alphabetSStart, alphabetSEnd, num, reverse)
		} else {
			s = c
		}
		shifted[i] = s
	}

	return string(shifted)
}

func shiftCore(cp rune, start rune, end rune, num int, reverse bool) rune {
	if reverse {
		s := cp - rune(num)
		if s < start {
			s = end - (start - s) + 1
		}

		return s
	}

	s := cp + rune(num)
	if s > end {
		s = start + (s - end) - 1
	}

	return s
}
