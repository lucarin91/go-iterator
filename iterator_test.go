package iterator

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/lucarin91/go-iterator/types"
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

	it := Map(ToIter(s), func(x int) string {
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

	got := Collect(ToIter(s))

	if !reflect.DeepEqual(got, s) {
		t.Errorf("got:%v, want:%v", got, s)
	}
}

func TestCollectOrError(t *testing.T) {
	s := []int{1, 2, 3, 4}

	// with no error
	it := Map(ToIter(s), func(x int) types.Result[int] {
		return types.Ok(x)
	})
	got, err := CollectWithError(it)
	if err != nil {
		t.Errorf("got err: %v, want nil error", err)
	}
	if !reflect.DeepEqual(got, s) {
		t.Errorf("got:%v, want:%v", got, s)
	}

	// with error
	it = Map(ToIter(s), func(x int) types.Result[int] {
		return types.Err[int](fmt.Errorf("this is an error"))
	})
	_, err = CollectWithError[int](it)
	if err == nil {
		t.Errorf("got nil error, want error")

	}
}

func TestCollectOrNone(t *testing.T) {
	// with all valid values
	s := []types.Option[int]{types.Some(1), types.Some(2), types.Some(3), types.Some(4)}
	got, ok := CollectWithOption(ToIter(s))
	if !ok {
		t.Error("got false, want true")
	}
	want := []int{1, 2, 3, 4}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got:%v, want:%v", got, want)
	}

	// with None
	s = []types.Option[int]{types.Some(1), types.None[int](), types.Some(3), types.Some(4)}
	_, ok = CollectWithOption[int](ToIter(s))
	if ok {
		t.Error("got true, want false")
	}
}

func TestFlatten(t *testing.T) {
	s := [][]int{{1, 2}, {3, 4}}

	it := Flatten(Map(ToIter(s), func(x []int) Iterator[int] {
		return ToIter(x)
	}))

	var got []int
	for it.Next() {
		got = append(got, it.Get())
	}

	want := []int{1, 2, 3, 4}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got:%v, want:%v", got, want)
	}
}

func TestFilter(t *testing.T) {
	s := []int{1, 2, 3, 4}

	it := Filter(ToIter(s), func(x int) bool {
		return x%2 == 0
	})

	var got []int
	for it.Next() {
		got = append(got, it.Get())
	}

	want := []int{2, 4}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got:%v, want:%v", got, want)
	}
}

func TestChunks(t *testing.T) {
	it := Chunks([]int{1, 2, 3, 4, 5}, 2)

	if !it.Next() {
		t.Errorf("should have next")
	}
	got, want := it.Get(), []int{1, 2}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got:%v, want:%v", got, want)
	}

	if !it.Next() {
		t.Errorf("should have next")
	}
	got, want = it.Get(), []int{3, 4}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got:%v, want:%v", got, want)
	}

	if !it.Next() {
		t.Errorf("should have next")
	}
	got, want = it.Get(), []int{5}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got:%v, want:%v", got, want)
	}

	if it.Next() {
		t.Errorf("should finish")
	}
}
