package main

import md "github.com/russross/blackfriday"

type Engine interface {
	Render(Post) []byte
}

func NewEngine() (engine Engine) {
	engine = new(defaultEngine)
	return
}

type defaultEngine struct{}

func (e *defaultEngine) Render(post Post) []byte {
	content := post.GetContent()

	if post.IsMarkdown() {
		return md.MarkdownCommon(content)
	}

	return content
}
