package p790

func numTilings(n int) int {
	mod := 1_000_000_007

	if n <= 2 {
		return n
	}

	fPrevious := 1
	fCurrent := 2
	pCurrent := 1

	for k := 3; k < n+1; k++ {
		tmp := fCurrent
		fCurrent = (fCurrent + fPrevious + 2*pCurrent) % mod
		pCurrent = (pCurrent + fPrevious) % mod
		fPrevious = tmp
	}

	return fCurrent
}
