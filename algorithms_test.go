package algorithms

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestBLnSrch(t *testing.T) {
	ns := []int{3, 2, 1, 0}
	ok, n := BLnSrch(
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

func TestRunConcurr(t *testing.T) {
	end := "end"
	fe := []func() error{
		func() (e error) {
			for {
			}
			return
		},
		func() error { return fmt.Errorf(end) },
	}
	e := RunConcurr(fe)
	require.Equal(t, end, e.Error())
}
