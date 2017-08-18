package popcount

// pc[i] is the population count of i.
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

func PopCountFor(x uint64) int {
	var sum int = 0
	for i := 0; i < 8; i++ {
		sum = sum + int(pc[byte(x>>(uint8(i)*8))])
	}
	return sum
}

func PopCountShift(x uint64) int {
	var sum int = 0
	if x == 0 {
		return 0
	} else {
		for {
			if x&1 == 1 {
				sum++
			}
			x = x >> 1
			if x == 0 {
				break
			}
		}
	}
	return sum
}

func PopCountUseBit(x uint64) int {
	var sum int = 0
	for {
		if x == 0 {
			break
		}
		x = x & (x - 1)
		sum++
	}
	return sum
}
