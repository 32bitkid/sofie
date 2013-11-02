package main_test

import (
	. "github.com/32bitkid/sofie"
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
