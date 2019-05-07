package algorithms

func Forall(inf func(int), n int) {
	for i := 0; i != n; i++ {
		inf(i)
	}
}

// BLnSrch is the Bounded Linear Search algorithm
func BLnSrch(ib func(int) bool, n int) (ok bool, i int) {
	i = 0
	for i != n && !ib(i) {
		i = i + 1
	}
	ok = i != n
	return
}

// TrueFF evaluates every function in fs until it it finds
// a false value returned by okf or evaluates all of them
func TrueFF(fs []func(), okf func() bool) (ok bool) {
	ib := func(i int) (b bool) {
		fs[i]()
		b = !okf()
		return
	}
	ok, _ = BLnSrch(ib, len(fs))
	return
}
