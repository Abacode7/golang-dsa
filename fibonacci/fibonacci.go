package fibonacci

/**
Todo: Fibonacci series:
	Fn = Fn-1 + Fn-2, where F0 = 0 and F1 = 1
*/

func fibN(n int) int {
	if n <= 0 {
		return 0
	}
	if n == 1 {
		return 1
	}

	fPrevPrev := 0
	fPrev := 1

	for i := 2; i <= n; i++ {
		temp := fPrev
		fPrev += fPrevPrev
		fPrevPrev = temp
	}

	return fPrev
}
