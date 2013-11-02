package main

import (
	md "github.com/russross/blackfriday"
	"io"
)

type Engine interface {
	Render(Post, io.Writer)
}

func NewEngine() (engine Engine) {
	engine = new(defaultEngine)
	return
}

type defaultEngine struct{}

func (e *defaultEngine) Render(post Post, w io.Writer) {
	content := post.GetContent()

	switch {
	case post.IsMarkdown():
		w.Write(md.MarkdownCommon(content))
	default:
		w.Write(content)
	}
}
