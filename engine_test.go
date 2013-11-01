package main_test

import (
  . "github.com/32bitkid/sofie"
	"bytes"
	"testing"
)

type mockPost struct {
	content []byte
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
