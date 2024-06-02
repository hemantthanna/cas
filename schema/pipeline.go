package schema

import "github.com/hemantthanna/cas/types"

func init() {
	registerCAS(&TransformOp{})
}

type TransformOp struct {
	Src types.Ref `json:"src"`
	Op  types.Ref `json:"op"`
	Dst types.Ref `json:"dst"`
}

func (t *TransformOp) References() []types.Ref {
	return []types.Ref{t.Src, t.Op, t.Dst}
}
