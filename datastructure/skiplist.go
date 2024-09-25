package datastructure

import (
	"cmp"
	"math/rand"
)

const (
	MaxLevel = 32
)

type SkipListNode[K cmp.Ordered, V any] struct {
	Key   K
	Value V
	Nexts [MaxLevel]*SkipListNode[K, V]
}

type SkipList[K cmp.Ordered, V any] struct {
	head  *SkipListNode[K, V]
	level int
}

func randomLevel() int {
	level := 1
	for ; level < MaxLevel; level++ {
		if rand.Intn(100) < 50 {
			break
		}
	}
	return level
}

func NewSkipList[K cmp.Ordered, V any]() *SkipList[K, V] {
	return &SkipList[K, V]{
		head:  &SkipListNode[K, V]{},
		level: 1,
	}
}

func (sl *SkipList[K, V]) Add(key K, value V) {
	level := randomLevel()
	if level > sl.level {
		sl.level = level
	}
	node := &SkipListNode[K, V]{Key: key, Value: value}
	curr := sl.head
	prevs := make([]*SkipListNode[K, V], level)
	for i := sl.level - 1; i >= 0; i-- {
		next := curr.Nexts[i]
		for next != nil && next.Key < key {
			curr = next
			next = curr.Nexts[i]
		}
		if i < level {
			prevs[i] = curr
		}
	}
	for i := 0; i < level; i++ {
		next := prevs[i].Nexts[i]
		node.Nexts[i] = next
		prevs[i].Nexts[i] = node
	}
}

func (sl *SkipList[K, V]) Find(key K) (V, bool) {
	curr := sl.head
	for i := sl.level - 1; i >= 0; i-- {
		// invariant: curr.Key < key
		next := curr.Nexts[i]
		for next != nil && next.Key < key {
			curr = next
			next = curr.Nexts[i]
		}
		// finally
		// next.Key >= key
		if next != nil && next.Key == key {
			return next.Value, true
		}
	}
	var v V
	return v, false
}

func (sl *SkipList[K, V]) Delete(key K) bool {
	curr := sl.head
	prevs := make([]*SkipListNode[K, V], sl.level)
	found := false
	for i := sl.level - 1; i >= 0; i-- {
		next := curr.Nexts[i]
		for next != nil && next.Key < key {
			curr = next
			next = curr.Nexts[i]
		}
		if next != nil && next.Key == key {
			prevs[i] = curr
			found = true
		}
	}
	for i, prev := range prevs {
		if prev != nil {
			prev.Nexts[i] = prev.Nexts[i].Nexts[i]
		}
	}
	return found
}
