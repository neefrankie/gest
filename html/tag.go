package html

import "strings"

type Tag struct {
	Token       Token
	SelfClosing bool
	Attr        Attr
	Children    []Node
}

func (t *Tag) leading() string {
	var out strings.Builder

	out.WriteByte('<')
	out.WriteString(string(t.Token))
	if !t.Attr.IsEmpty() {
		out.WriteByte(' ')
		out.WriteString(t.Attr.Encode())
	}
	if t.SelfClosing {
		out.WriteString("/>")
	} else {
		out.WriteByte('>')
	}

	return out.String()
}

func (t *Tag) trailing() string {
	if t.SelfClosing {
		return ""
	}
	return "</" + string(t.Token) + ">"
}

func (t *Tag) TokenLiteral() string {
	return string(t.Token)
}

func (t *Tag) String() string {
	var out strings.Builder

	out.WriteString(t.leading())

	for _, node := range t.Children {
		out.WriteString(node.String())
	}

	out.WriteString(t.trailing())

	return out.String()
}
