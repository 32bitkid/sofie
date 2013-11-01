package main

type Engine interface {
	Render(Post) string
}

func NewEngine() (engine Engine) {
	engine = new(defaultEngine)
	return
}

type defaultEngine struct{}

func (e *defaultEngine) Render(post Post) string {
	return post.GetContent()
}
