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

	it := Map[int, string](ToIter(s), func(x int) string {
		return fmt.Sprintf("n:%d", x)
	})

	var got []string
	for it.Next() {
		got = append(got, it.Get())
	}

	want := []string{"n:1", "n:2", "n:3", "n:4"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got:%v, want:%v", got, want)
	}
}

func TestCollect(t *testing.T) {
	s := []int{1, 2, 3, 4}

	got := Collect[int](ToIter(s))

	if !reflect.DeepEqual(got, s) {
		t.Errorf("got:%v, want:%v", got, s)
	}
}

func TestFlatten(t *testing.T) {
	s := [][]int{{1, 2}, {3, 4}}

	it := Map[[]int, Iterator[int]](ToIter(s), func(x []int) Iterator[int] {
		return ToIter(x)
	})

	it3 := Flatten(it)

	var got []int
	for it3.Next() {
		got = append(got, it3.Get())
	}

	want := []int{1, 2, 3, 4}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got:%v, want:%v", got, want)
	}
}
