package store

type (
	Pred interface {
		AcceptPred(v PredVisitor)
		prec() uint8
	}

	PredVisitor interface {
		VisitOr(p OrPred)
		VisitAnd(p AndPred)
		VisitNot(p NotPred)
		VisitBin(p BinPred)
	}

	OrPred struct {
		ops []Pred
	}

	AndPred struct {
		ops []Pred
	}

	NotPred struct {
		op Pred
	}

	BinPred struct {
		op    BinOp
		path  Path
		value any
	}

	BinOp uint8
)

func Or(ops ...Pred) Pred {
	items := make([]Pred, 0, len(ops))

	for _, op := range ops {
		if top, ok := op.(OrPred); ok {
			items = append(items, top.ops...)
		} else if op != nil {
			items = append(items, op)
		}
	}

	switch len(items) {
	case 0:
		return nil
	case 1:
		return items[0]
	}

	return OrPred{ops: items}
}

func And(ops ...Pred) Pred {
	items := make([]Pred, 0, len(ops))

	for _, op := range ops {
		if top, ok := op.(AndPred); ok {
			items = append(items, top.ops...)
		} else if op != nil {
			items = append(items, op)
		}
	}

	switch len(items) {
	case 0:
		return nil
	case 1:
		return items[0]
	}

	return AndPred{ops: items}
}

func Not(op Pred) Pred {
	if top, ok := op.(NotPred); ok {
		return top.op
	} else if op == nil {
		return nil
	}

	return NotPred{op: op}
}

func Eq(path Path, value any) BinPred {
	return BinPred{
		op:    BinEq,
		path:  path,
		value: value,
	}
}

func Neq(path Path, value any) BinPred {
	return BinPred{
		op:    BinNeq,
		path:  path,
		value: value,
	}
}

func Lt(path Path, value any) BinPred {
	return BinPred{
		op:    BinLt,
		path:  path,
		value: value,
	}
}

func Leq(path Path, value any) BinPred {
	return BinPred{
		op:    BinLeq,
		path:  path,
		value: value,
	}
}

func Gt(path Path, value any) BinPred {
	return BinPred{
		op:    BinGt,
		path:  path,
		value: value,
	}
}

func Geq(path Path, value any) BinPred {
	return BinPred{
		op:    BinGeq,
		path:  path,
		value: value,
	}
}

func Like(path Path, value any) BinPred {
	return BinPred{
		op:    BinLike,
		path:  path,
		value: value,
	}
}

const (
	BinEq BinOp = iota
	BinNeq
	BinLt
	BinLeq
	BinGt
	BinGeq
	BinLike
)

func (p OrPred) Ops() []Pred  { return p.ops }
func (p AndPred) Ops() []Pred { return p.ops }
func (p NotPred) Op() Pred    { return p.op }
func (p BinPred) Op() BinOp   { return p.op }
func (p BinPred) Path() Path  { return p.path }
func (p BinPred) Value() any  { return p.value }

func (p OrPred) AcceptPred(v PredVisitor)  { v.VisitOr(p) }
func (p AndPred) AcceptPred(v PredVisitor) { v.VisitAnd(p) }
func (p NotPred) AcceptPred(v PredVisitor) { v.VisitNot(p) }
func (p BinPred) AcceptPred(v PredVisitor) { v.VisitBin(p) }

func (OrPred) prec() uint8  { return 0 }
func (AndPred) prec() uint8 { return 1 }
func (NotPred) prec() uint8 { return 2 }
func (BinPred) prec() uint8 { return 3 }

var (
	_ Pred = OrPred{}
	_ Pred = AndPred{}
	_ Pred = NotPred{}
	_ Pred = BinPred{}
)
