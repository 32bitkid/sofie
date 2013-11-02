package main

import (
	md "github.com/russross/blackfriday"
	"io"
)

type Engine interface {
	Render(Post, io.Writer)
	SetDefaultLayout(Layout)
}

func NewEngine() (engine Engine) {
	engine = new(defaultEngine)
	return
}

type defaultEngine struct {
	defaultLayout Layout
}

func (e *defaultEngine) Render(post Post, w io.Writer) {
	content := post.GetContent()
	layout := e.defaultLayout

	if layout != nil {
		layout.RenderBefore(w)
	}

	switch {
	case post.IsMarkdown():
		w.Write(md.MarkdownCommon(content))
	default:
		w.Write(content)
	}

	if layout != nil {
		layout.RenderAfter(w)
	}
}

func (e *defaultEngine) SetDefaultLayout(layout Layout) {
	e.defaultLayout = layout
}
