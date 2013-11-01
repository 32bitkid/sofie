package main

import "testing"

type mockPost struct {
	content string
}

func (m *mockPost) GetContent() string {
	return m.content
}

var htmlPost = &mockPost{"<h1>This is a test</h1>"}

func TestRenderingHtmlPost(t *testing.T) {
	engine := NewEngine()
	actual := engine.Render(htmlPost)
	expected := "<h1>This is a test</h1>"
	if actual != expected {
		t.Errorf("Expected %q, got %q", expected, actual)
	}
}
