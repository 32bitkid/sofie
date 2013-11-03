package main_test

import (
	. "github.com/32bitkid/sofie"
	"io/ioutil"
	"os"
	"testing"
)

func TestParsingFromYaml(t *testing.T) {
	yaml := []byte(`content: here is some content`)

	_, err := PostFromYaml(yaml)

	if err != nil {
		t.Error(err)
	}
}

func TestInvalidInput(t *testing.T) {
	yaml := []byte(`!garbage@<this>is<garbage>`)

	_, err := PostFromYaml(yaml)

	if err == nil {
		t.Error(err)
	}
}

func ExamplePost_Render_rawContent() {
	yaml := []byte(`content: here is some content`)

	post, _ := PostFromYaml(yaml)

	formatters := FormatterList{
		"": NoopFormatter,
	}

	post.Render(os.Stdout, formatters)
	// Output:
	// here is some content
}

func TestUsingADifferentFormatter(t *testing.T) {
	yaml := []byte(`
content: here is some content
format: custom
`)

	post, _ := PostFromYaml(yaml)

	formatterCalled := false

	formatters := FormatterList{
		"custom": func(data []byte) []byte {
			formatterCalled = true
			return data
		},
	}

	post.Render(ioutil.Discard, formatters)

	if formatterCalled != true {
		t.Error("Specified formatter was not invoked!")
	}
}
