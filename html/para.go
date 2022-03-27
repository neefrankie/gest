package html

type ParaParams struct {
	ClassName string
	Children  []Node
}

type Para struct {
	ClassName string
	Children  []Node
}

func (p *Para) blockNode() {}
