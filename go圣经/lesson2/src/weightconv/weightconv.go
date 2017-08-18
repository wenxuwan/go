package weightconv

func KToP(k Kilogram) Pound {
	return Pound(k / 0.4535924)
}

func PToK(p Pound) Kilogram {
	return Kilogram(0.4535924 * p)
}
