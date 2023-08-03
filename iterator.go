package iterator

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
