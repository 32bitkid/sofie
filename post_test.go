package main_test

import (
  . "github.com/32bitkid/sofie"
  "testing"
  "bytes"
  )

type postFn func(*testing.T) Post


func TestGettingYamlContentFromPost(t *testing.T) {
  
  yaml := []byte(`content: here is some content`)
  
  post, err := PostFromYaml(yaml)
  
  if err != nil {
    t.Error(err)
    return 
  }
  
  actual := post.GetContent()
  expected := []byte(`here is some content`)
  
  if !bytes.Equal(actual, expected) {
    t.Errorf("Expected %q, but got %q", expected, actual)
  }
}

func TestUnspecifiedFormatOfPost(t *testing.T) {
  yaml := []byte(`content: here is some content`)
  
  post, err := PostFromYaml(yaml)
  
  if err != nil {
    t.Error(err)
  }
  
  if post.IsMarkdown() {
    t.Error("Expected to not be markdown")
  }
}

func TestPostIsMarkdown(t *testing.T) {
  yaml := []byte(`
content: this is markdown
format: md
`)
  
  post, err := PostFromYaml(yaml)
  
  if err != nil {
    t.Error(err)
    return 
  }
 
  if !post.IsMarkdown() {
    t.Error("Expected post to be formatted with markdown")
  }
  
}