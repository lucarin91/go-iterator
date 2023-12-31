package iterator

import "github.com/lucarin91/go-iterator/types"

type Iterator[T any] interface {
	Next() bool
	Get() T
}

type SliceIt[T any] struct {
	inner []T
	idx   int
}

func ToIter[T any](slice []T) Iterator[T] {
	return &SliceIt[T]{
		inner: slice,
		idx:   -1,
	}
}

func (s *SliceIt[T]) Next() bool {
	s.idx += 1
	return s.idx < len(s.inner)
}

func (s *SliceIt[T]) Get() T {
	return s.inner[s.idx]
}

func Map[T1 any, T2 any](it Iterator[T1], f func(T1) T2) Iterator[T2] {
	return &MapIt[T1, T2]{
		it: it,
		f:  f,
	}
}

type MapIt[T1 any, T2 any] struct {
	it Iterator[T1]
	f  func(T1) T2
}

func (s *MapIt[T1, T2]) Next() bool {
	return s.it.Next()
}

func (s *MapIt[T1, T2]) Get() T2 {
	return s.f(s.it.Get())
}

func Collect[T any](it Iterator[T]) []T {
	var out []T
	for it.Next() {
		out = append(out, it.Get())
	}
	return out
}

func CollectWithError[T any](it Iterator[types.Result[T]]) ([]T, error) {
	var out []T
	for it.Next() {
		value, err := it.Get().Unwrap()
		if err != nil {
			return out, err
		}
		out = append(out, value)
	}
	return out, nil
}

func CollectWithOption[T any](it Iterator[types.Option[T]]) ([]T, bool) {
	var out []T
	for it.Next() {
		value, ok := it.Get().Unwrap()
		if !ok {
			return out, ok
		}
		out = append(out, value)
	}
	return out, true
}

func Flatten[T any](it Iterator[Iterator[T]]) Iterator[T] {
	return &FlattenIt[T]{
		it: it,
	}
}

type FlattenIt[T any] struct {
	it    Iterator[Iterator[T]]
	inner Iterator[T]
}

func (s *FlattenIt[T]) Next() bool {
	if s.inner == nil {
		if s.it.Next() {
			s.inner = s.it.Get()
			return s.Next()
		}
		return false
	}
	if s.inner.Next() {
		return true
	}
	s.inner = nil
	return s.Next()
}

func (s *FlattenIt[T]) Get() T {
	return s.inner.Get()
}

func Filter[T any](it Iterator[T], f func(x T) bool) Iterator[T] {
	return &FilterIt[T]{
		it: it,
		f:  f,
	}
}

type FilterIt[T any] struct {
	it Iterator[T]
	f  func(x T) bool
}

func (s *FilterIt[T]) Next() bool {
	for s.it.Next() {
		if s.f(s.it.Get()) {
			return true
		}
	}
	return false
}

func (s *FilterIt[T]) Get() T {
	return s.it.Get()
}

func Chunks[T any](input []T, size int) Iterator[[]T] {
	return &ChunksIt[T]{
		p:     0,
		input: input,
		size:  size,
	}
}

type ChunksIt[T any] struct {
	p     int
	input []T
	size  int
	next  []T
}

func (b *ChunksIt[T]) Next() bool {
	if b.p >= len(b.input) {
		return false
	}

	start := b.p
	end := min(b.p+b.size, len(b.input))

	b.p = end
	b.next = b.input[start:end]

	return true
}

func (b *ChunksIt[T]) Get() []T {
	return b.next
}
