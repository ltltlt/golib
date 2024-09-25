package datastructure

import (
	"cmp"
	"slices"
)

type OrderedSet[K cmp.Ordered, V any] interface {
	Add(key K, value V)
	Find(key K) (V, bool)
	Delete(key K) bool
}

type Array[K cmp.Ordered, V any] struct {
	slice []pair[K, V]
}

type pair[K cmp.Ordered, V any] struct {
	key   K
	value V
}

func NewArray[K cmp.Ordered, V any]() *Array[K, V] {
	return &Array[K, V]{}
}

func (arr *Array[K, V]) Add(key K, value V) {
	idx, _ := slices.BinarySearchFunc(arr.slice, key, func(p1 pair[K, V], k K) int {
		return cmp.Compare(p1.key, k)
	})
	arr.slice = slices.Insert(arr.slice, idx, pair[K, V]{key: key, value: value})
}

func (arr *Array[K, V]) Find(key K) (V, bool) {
	idx, ok := slices.BinarySearchFunc(arr.slice, key, func(p1 pair[K, V], k K) int {
		return cmp.Compare(p1.key, k)
	})
	var v V
	if ok {
		v = arr.slice[idx].value
	}
	return v, ok
}

func (arr *Array[K, V]) Delete(key K) bool {
	v := slices.DeleteFunc(arr.slice, func(p pair[K, V]) bool {
		return p.key == key
	})
	return len(v) > 0
}
