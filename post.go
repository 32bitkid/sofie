package main

import "launchpad.net/goyaml"

type Post interface {
	GetContent() string
}

type yamlPost struct {
	Content string
}

func (p *yamlPost) GetContent() string {
	return p.Content
}

func PostFromYaml(yaml []byte) (post Post, err error) {
	post = new(yamlPost)
	err = goyaml.Unmarshal(yaml, post)
	return
}
