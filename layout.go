package main

import "io"

type Layout interface {
	RenderBefore(io.Writer)
	RenderAfter(io.Writer)
}

type EmptyLayout struct{}

func (l *EmptyLayout) RenderBefore(w io.Writer) {}
func (l *EmptyLayout) RenderAfter(w io.Writer)  {}
