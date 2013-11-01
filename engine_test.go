package main

import "testing"

var htmlPost = []byte(`
content: <h1>This is a test<h1>
`)

func postWrapper(yaml []byte, t *testing.T) (post Post) {
	post, err := PostFromYaml(yaml)

	if err != nil {
		t.Fail()
	}

	return
}

type PostFn func(*testing.T) Post

var testPost PostFn = func(t *testing.T) Post { return postWrapper(htmlPost, t) }

func TestRenderingHtmlPost(t *testing.T) {
	engine := NewEngine()
	actual := engine.Render(testPost(t))
	expected := "<h1>This is a test<h1>"
	if actual != expected {
		t.Errorf("Expected %q, got %q", expected, actual)
	}
}
