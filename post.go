package main

import "launchpad.net/goyaml"

type Post interface {
	GetContent() []byte
	IsMarkdown() bool
}

type yamlPost struct {
	Content []byte
}

func (p *yamlPost) GetContent() []byte {
	return p.Content
}

func (p *yamlPost) IsMarkdown() bool {
	return false
}

func PostFromYaml(yaml []byte) (post Post, err error) {
	post = new(yamlPost)
	err = goyaml.Unmarshal(yaml, post)
	return
}
