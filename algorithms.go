package algorithms

// Forall calls the function with all integers between 0 and n
// (upper bound is excluded), in increasing order
func Forall(inf func(int), n int) {
	for i := 0; i != n; i++ {
		inf(i)
	}
}

// BLS is the Bounded Linear Search algorithm
// https://www.cs.utexas.edu/users/EWD/transcriptions/EWD09xx/EWD930.html
func BLS(ib func(int) bool, n int) (ok bool, i int) {
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
	foundFalse, _ := BLS(ib, len(fs))
	ok = !foundFalse
	return
}

// KFunc is a string-function pair
type KFunc struct {
	Key  string
	Func func()
}

// ExecF executes the corresponding function for the supplied
// key
func ExecF(kf []KFunc, key string) (ok bool) {
	ib := func(i int) bool { return kf[i].Key == key }
	ok, n := BLS(ib, len(kf))
	if ok {
		kf[n].Func()
	}
	return
}

// RunConcurrent runs concurrently all functions until one
// of them returns an error
func RunConcurrent(fe []func() error) (e error) {
	ec := make(chan error)
	f := func(i int) { go func() { ec <- fe[i]() }() }
	Forall(f, len(fe))
	e = <-ec
	return
}

// Max returns the maximum between two integers
func Max(a, b int) (c int) {
	if b > a {
		c = b
	} else {
		c = a
	}
	return
}

// Min returns the mininum between two integers
func Min(a, b int) (c int) {
	if a < b {
		c = a
	} else {
		c = b
	}
	return
}

// Reverse calls the function with all the values between 0
// and n (excluded upper bound), in decreasing order
func Reverse(fi func(int), n int) {
	for i := 0; i != n; i++ {
		fi(n - i - 1)
	}
}
