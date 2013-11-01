package main

import "launchpad.net/goyaml"

var markdownFormats = [...]string{"md", "mdown", "markdown"}

type Post interface {
	GetContent() []byte
	IsMarkdown() bool
}

type yamlPost struct {
	Content string
	Format  string
}

func (p *yamlPost) GetContent() []byte {
	return []byte(p.Content)
}

func (p *yamlPost) IsMarkdown() bool {
	for _, ext := range markdownFormats {
		if ext == p.Format {
			return true
		}
	}
	return false
}

func PostFromYaml(yaml []byte) (post Post, err error) {
	post = new(yamlPost)
	err = goyaml.Unmarshal(yaml, post)
	return
}
