package goapi

import (
	"testing"

	"github.com/ysmood/got"
)

func TestPath(t *testing.T) {
	g := got.T(t)

	p, err := newPath("/a/b/c")
	g.E(err)

	g.Nil(p.match("/x"))
	g.Eq(p.match("/a/b/c"), map[string]string{})

	p, err = newPath("/a/{b}/c/{d}")
	g.E(err)

	g.Eq(p.match("/a/x/c/y"), map[string]string{"b": "x", "d": "y"})
}
