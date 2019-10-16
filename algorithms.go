package algorithms

func Forall(inf func(int), n int) {
	for i := 0; i != n; i++ {
		inf(i)
	}
}

// BLnSrch is the Bounded Linear Search algorithm
// https://www.cs.utexas.edu/users/EWD/transcriptions/EWD09xx/EWD930.html
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
	foundFalse, _ := BLnSrch(ib, len(fs))
	ok = !foundFalse
	return
}

type KFunc struct {
	Key  string
	Func func()
}

func ExecF(kf []KFunc, key string) (ok bool) {
	ib := func(i int) bool { return kf[i].Key == key }
	ok, n := BLnSrch(ib, len(kf))
	if ok {
		kf[n].Func()
	}
	return
}

func RunConcurrent(fe []func() error) (e error) {
	ec := make(chan error)
	f := func(i int) { go func() { ec <- fe[i]() }() }
	Forall(f, len(fe))
	e = <-ec
	return
}

func Max(a, b int) (c int) {
	if b > a {
		c = b
	} else {
		c = a
	}
	return
}

func Min(a, b int) (c int) {
	if a < b {
		c = a
	} else {
		c = b
	}
	return
}

func Reverse(fi func(int), n int) {
	for i := 0; i != n; i++ {
		fi(n - i - 1)
	}
}
