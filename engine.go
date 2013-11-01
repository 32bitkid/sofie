package main

type Engine interface{}

func NewEngine() (engine Engine) {
	engine = new(defaultEngine)
	return
}

type defaultEngine struct{}
