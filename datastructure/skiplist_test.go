package datastructure

import (
	"math/rand"
	"testing"
)

func TestSkipList1(t *testing.T) {
	sl := NewSkipList[int, int]()
	for i := 0; i < 5; i++ {
		sl.Add(i, i)
	}
	for i := 0; i < 5; i++ {
		v, ok := sl.Find(i)
		if !ok || v != i {
			t.Fatalf("expected k %v, v %v, got %v, %v\n", i, i, v, ok)
		}
	}
	if sl.Delete(5) {
		t.Fatalf("delete 5 success while 5 is not added")
	}
	for i := 0; i < 3; i++ {
		if !sl.Delete(i) {
			t.Fatalf("failed to delete k %v", i)
		}
	}
	if v, ok := sl.Find(3); v != 3 || !ok {
		t.Fatalf("expect k %v, v %v after delete, got %v, %v\n", 3, 3, v, ok)
	}
	if v, ok := sl.Find(4); v != 4 || !ok {
		t.Fatalf("expect k %v, v %v after delete, got %v, %v\n", 4, 4, v, ok)
	}
}

func TestSkipList2(t *testing.T) {
	sl := NewSkipList[int, int]()
	doOp(t, sl)
}

func BenchmarkArray(b *testing.B) {
	arr := NewArray[int, int]()
	doOp(b, arr)
}

func BenchmarkSkipList(b *testing.B) {
	arr := NewSkipList[int, int]()
	doOp(b, arr)
}

func doOp(t testing.TB, sl OrderedSet[int, int]) {
	added := map[int]int{}
	addCount := 50000
	deleteCount := 500
	for i := 0; i < addCount; i++ {
		var k, v int
		for {
			k = rand.Int()
			v = rand.Int()
			if _, ok := added[k]; !ok {
				added[k] = v
				break
			}
		}
		sl.Add(k, v)
	}
	for k, v := range added {
		if v1, ok := sl.Find(k); !ok || v1 != v {
			t.Fatalf("expected k %v, v %v, got %v, %v\n", k, v, v1, ok)
		}
	}
	for i := 0; i < deleteCount; i++ {
		var k int
		for k = range added {
			break
		}
		delete(added, k)
		if !sl.Delete(k) {
			t.Fatalf("failed to delete k %v\n", k)
		}
		if _, ok := sl.Find(k); ok {
			t.Fatalf("after delete of k %v, still exists\n", k)
		}
	}
}
