package store

type (
	Path []PathItem

	PathItem interface {
		AcceptPath(v PathVisitor)
	}

	PathVisitor interface {
		VisitField(f PathField)
		VisitIndex(f PathIndex)
	}

	PathField string
	PathIndex int64
)

func (f PathField) AcceptPath(v PathVisitor) { v.VisitField(f) }
func (i PathIndex) AcceptPath(v PathVisitor) { v.VisitIndex(i) }

var (
	_ PathItem = PathField("")
	_ PathItem = PathIndex(0)
)
