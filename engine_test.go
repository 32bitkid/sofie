package main_test

import (
	"bytes"
	. "github.com/32bitkid/sofie"
	"testing"
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

func TestRenderingRawPost(t *testing.T) {
	var buffer bytes.Buffer
	NewEngine().Render(rawPost, &buffer)

	actual := buffer.Bytes()
	expected := []byte("<h1>This is a test</h1>")
	if !bytes.Equal(actual, expected) {
		t.Errorf("Expected %q, got %q", expected, actual)
	}
}

func TestRenderingMarkdownPost(t *testing.T) {
	var buffer bytes.Buffer

	NewEngine().Render(mdPost, &buffer)
	actual := buffer.Bytes()

	expected := []byte("<h1>This is a test</h1>\n")
	if !bytes.Equal(actual, expected) {
		t.Errorf("Expected %q, got %q", expected, actual)
	}
}
