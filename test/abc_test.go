package test

import (
	"github.com/stretchr/testify/require"
	"testing"
)

type foo[T any] struct {
	a *T
}

func (f *foo[T]) name(e *T) {

}

func Test(t *testing.T) {
	f := foo[int]{}
	a := 1
	f.name(&a)
	require.NoError(t, nil)
}
