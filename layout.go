package main

import "io"

type Layout interface {
	RenderBefore(io.Writer)
	RenderAfter(io.Writer)
}