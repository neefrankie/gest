package html

import "strings"

type Params struct {
	Lang string
	Head *Head
	Body *Body
}

type HTML struct {
	Tag
}

func NewHTML(p Params) *HTML {
	return &HTML{
		Tag{
			Token:       TokenHTML,
			SelfClosing: false,
			Attr:        NewAttr().Set("lang", p.Lang),
			Children: []Node{
				p.Head,
				p.Body,
			},
		},
	}
}

func (h *HTML) String() string {
	var out strings.Builder
	out.WriteString("<!DOCTYPE html>")

	out.WriteString(h.Tag.String())

	return out.String()
}
