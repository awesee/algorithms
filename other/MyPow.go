package other

/*
Implement pow(x, n).
*/

func MyPow(x float64, n int) float64 {

	r := 1.0

	if n < 0 {
		n = -n
		x = 1 / x
	}

	for ; n > 0; n >>= 1 {
		if n&1 == 1 {
			r *= x
		}
		x *= x
	}

	return r
}

func MyPowV1(x float64, n int) float64 {

	r := 1.0
	if n == 0 {
		return r
	}

	if n < 0 {
		n = -n
		x = 1 / x
	}

	if n&1 == 0 {
		r = MyPow(x*x, n>>1)
	} else {
		r = MyPow(x*x, n>>1) * x
	}

	return r
}
