package main

import (
	"io"
)

type Engine interface {
	Render(Post, io.Writer)
	SetDefaultLayout(Layout)
}

func NewEngine() Engine {
	return &defaultEngine{
		defaultLayout: new(EmptyLayout),
		formatters:    DefaultFormatters(),
	}
}

type defaultEngine struct {
	defaultLayout Layout
	formatters    Formatters
}

func (e *defaultEngine) Render(post Post, w io.Writer) {
	layout := e.defaultLayout

	layout.RenderBefore(w)
	post.Render(w, e.formatters)
	layout.RenderAfter(w)
}

func (e *defaultEngine) SetDefaultLayout(layout Layout) {
	e.defaultLayout = layout
}
