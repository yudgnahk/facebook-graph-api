package utils

import (
	"golang.org/x/exp/constraints"
)

// PointerAble define an interface that can change to pointer, include string, int, int64, unit
type PointerAble interface {
	constraints.Integer | constraints.Unsigned | string
}

func GetPointer[T PointerAble](value T) *T {
	return &value
}
