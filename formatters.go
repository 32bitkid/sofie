package main

import md "github.com/russross/blackfriday"

type Formatters interface {
	Fetch(string) ContentFormatter
}

type (
	ContentFormatter func([]byte) []byte
	FormatterList    map[string]ContentFormatter
)

var NoopFormatter ContentFormatter = func(input []byte) []byte {
	return input
}

func DefaultFormatters() Formatters {
	return FormatterList{
		"":         NoopFormatter,
		"raw":      NoopFormatter,
		"md":       md.MarkdownCommon,
		"mdown":    md.MarkdownCommon,
		"markdown": md.MarkdownCommon,
	}
}

func (list FormatterList) Fetch(name string) ContentFormatter {
	formatter, found := list[name]
	if !found {
		panic("Formatter not found")
	}
	return formatter
}
