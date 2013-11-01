package main

import (
	"bytes"
	"testing"
)

type mockPost struct {
	content []byte
	format  string
}

func (m *mockPost) GetContent() []byte {
	return m.content
}

func (m *mockPost) IsMarkdown() bool {
	return m.format == "md"
}

var rawPost = &mockPost{[]byte("<h1>This is a test</h1>"), "raw"}
var mdPost = &mockPost{[]byte("# This is a test"), "md"}

func TestRenderingRawPost(t *testing.T) {
	engine := NewEngine()
	actual := engine.Render(rawPost)
	expected := []byte("<h1>This is a test</h1>")
	if !bytes.Equal(actual, expected) {
		t.Errorf("Expected %q, got %q", expected, actual)
	}
}

func TestRenderingMarkdownPost(t *testing.T) {
	engine := NewEngine()
	actual := engine.Render(mdPost)
	expected := []byte("<h1>This is a test</h1>\n")
	if !bytes.Equal(actual, expected) {
		t.Errorf("Expected %q, got %q", expected, actual)
	}
}
