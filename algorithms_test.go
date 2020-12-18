package algorithms

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBLS(t *testing.T) {
	ns := []int{3, 2, 1, 0}
	ok, n := BLS(
		func(i int) bool { return ns[i]%2 == 0 },
		len(ns),
	)
	require.True(t, ok)
	require.Equal(t, 1, n)
}

func TestTrueFF(t *testing.T) {
	v := true
	fs := []func(){
		func() { v = true },
		func() { v = false },
	}
	ok := TrueFF(fs, func() bool { return v })
	require.False(t, ok)
}

func TestExecF(t *testing.T) {
	hello, world := "hola", "mundo"
	v := false
	kf := []KFunc{
		{
			hello,
			func() {},
		},
		{
			world,
			func() { v = true },
		},
	}
	ok := ExecF(kf, world)
	require.True(t, ok)
	require.True(t, v)
}

func TestRunConcurrent(t *testing.T) {
	end := "end"
	fe := []func() error{
		func() (e error) {
			for {
			}
		},
		func() error { return fmt.Errorf(end) },
	}
	e := RunConcurrent(fe)
	require.Equal(t, end, e.Error())
}

func TestMax(t *testing.T) {
	m, n := Max(0, 0), Max(-1, 1)
	require.Equal(t, 0, m)
	require.Equal(t, 1, n)
}

func TestMin(t *testing.T) {
	m, n := Min(0, 0), Min(-1, 1)
	require.Equal(t, 0, m)
	require.Equal(t, -1, n)
}

func TestReverse(t *testing.T) {
	s, n := make([]int, 10), 0
	f := func(i int) { s[n], n = i, n+1 }
	Reverse(f, len(s))
	require.Equal(t, []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}, s)
}
