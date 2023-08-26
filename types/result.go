package types

type Result[T any] struct {
	value T
	err   error
}

func Err[T any](err error) Result[T] {
	return Result[T]{err: err}
}

func Ok[T any](value T) Result[T] {
	return Result[T]{value: value}
}

func (r Result[T]) Unwrap() (T, error) {
	return r.value, r.err
}

type ResultIt[T any] struct {
	ended  bool
	result Result[T]
}

func (r Result[T]) ToIter() *ResultIt[T] {
	return &ResultIt[T]{
		ended:  false,
		result: r,
	}
}

func (s *ResultIt[T]) Next() bool {
	if s.ended {
		return false
	}
	s.ended = true
	return s.result.err == nil
}

func (s *ResultIt[T]) Get() T {
	return s.result.value
}
