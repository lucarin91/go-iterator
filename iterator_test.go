package iterator

import (
	"fmt"
	"reflect"
	"testing"
)

func TestIterators(t *testing.T) {
	s := []int{1, 2, 3, 4}
	it := ToIter(s)

	var got []int
	for it.Next() {
		got = append(got, it.Get())
	}

	if !reflect.DeepEqual(got, s) {
		t.Errorf("got:%v, want:%v", got, s)
	}
}

func TestMap(t *testing.T) {
	s := []int{1, 2, 3, 4}
	it := ToIter(s)

	it2 := Map[int, string](&it, func(x int) string {
		return fmt.Sprintf("n:%d", x)
	})

	var got []string
	for it2.Next() {
		got = append(got, it2.Get())
	}

	want := []string{"n:1", "n:2", "n:3", "n:4"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got:%v, want:%v", got, want)
	}
}

func TestCollect(t *testing.T) {
	s := []int{1, 2, 3, 4}

	it := ToIter(s)

	got := Collect[int](&it)

	if !reflect.DeepEqual(got, s) {
		t.Errorf("got:%v, want:%v", got, s)
	}
}

func TestFlatten(t *testing.T) {
	s := [][]int{{1, 2}, {3, 4}}

	it := ToIter(s)

	it2 := Map[[]int, Iterator[int]](&it, func(x []int) Iterator[int] {
		it := ToIter(x)
		return &it
	})

	it3 := Flatten(it2)

	var got []int
	for it3.Next() {
		got = append(got, it3.Get())
	}

	want := []int{1, 2, 3, 4}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got:%v, want:%v", got, want)
	}
}
