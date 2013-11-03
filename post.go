package main

import (
	"io"
	"launchpad.net/goyaml"
)

type Post interface {
	Render(io.Writer, Formatters)
}

type yamlPost struct {
	Content string
	Format  string
}

func (p *yamlPost) Render(w io.Writer, f Formatters) {
	w.Write(f.Fetch(p.Format)([]byte(p.Content)))
}

func PostFromYaml(yaml []byte) (post Post, err error) {
	post = new(yamlPost)
	err = goyaml.Unmarshal(yaml, post)
	return
}
