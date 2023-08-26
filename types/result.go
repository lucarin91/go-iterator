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

