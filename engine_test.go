package main_test

import (
	. "github.com/32bitkid/sofie"
	"io"
	"os"
)

type mockPost struct {
	content    []byte
	isMarkdown bool
}

func (m *mockPost) GetContent() []byte {
	return m.content
}

func (m *mockPost) IsMarkdown() bool {
	return m.isMarkdown
}

var rawPost = &mockPost{[]byte("<h1>This is a test</h1>"), false}
var mdPost = &mockPost{[]byte("# This is a test"), true}

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
