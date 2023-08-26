package types

type Option[T any] struct {
	value   T
	present bool
}

func Some[T any](value T) Option[T] {
	return Option[T]{value: value, present: true}
}

func None[T any]() Option[T] {
	return Option[T]{present: false}
}

func (o Option[T]) Unwrap() (T, bool) {
	return o.value, o.present
}

func (o Option[T]) Get() T {
	if !o.present {
		panic("option value is None")
	}
	return o.value
}

func (o Option[_]) IsValid() bool {
	return o.present
}
