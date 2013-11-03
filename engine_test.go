package main_test

import (
	. "github.com/32bitkid/sofie"
	"io"
	"os"
)

type mockPost struct {
	content []byte
	format  string
}

func (m *mockPost) Render(w io.Writer, f Formatters) {
	w.Write(f.Fetch(m.format)(m.content))
}

var rawPost = &mockPost{[]byte("<h1>This is a test</h1>"), ""}
var mdPost = &mockPost{[]byte("# This is a test"), "md"}

func ExampleEngine_Render_rawFormat() {
	NewEngine().Render(rawPost, os.Stdout)
	// Output:
	// <h1>This is a test</h1>
}

func ExampleEngine_Render_markdownFormat() {
	NewEngine().Render(mdPost, os.Stdout)
	// Output:
	// <h1>This is a test</h1>
}

type mockLayout struct{}

func (l *mockLayout) RenderBefore(w io.Writer) {
	w.Write([]byte("before"))
}

func (l *mockLayout) RenderAfter(w io.Writer) {
	w.Write([]byte("after"))
}

func ExampleEngine_Render_withDefaultLayout() {
	engine := NewEngine()
	engine.SetDefaultLayout(new(mockLayout))
	engine.Render(rawPost, os.Stdout)
	// Output:
	// before<h1>This is a test</h1>after
}
