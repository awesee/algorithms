package other

func ReverseInt(x int) int {

	res, t := 0, 0
	for x != 0 {
		t = res*10 + x%10
		if t/10 != res {
			return 0
		}
		res = t
		x /= 10
	}

	return res
}
