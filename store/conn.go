package store

import "context"

type (
	Conn interface {
		Context() context.Context
		Exists(pred Pred) (bool, error)
		Count(pred Pred) (uint64, error)
		First(pred Pred) (*Object, error)
		Distinct(path Path, pred Pred) ([]any, error)
		All(pred Pred, order Ordering, offset uint64, limit uint64) ([]Object, error)
	}

	Tx interface {
		Conn
		Create(obj *Object) error
		Update(obj *Object) error
		Delete(obj Object) error
	}

	Ordering []Order

	Order struct {
		Path Path
		Desc bool
	}
)
