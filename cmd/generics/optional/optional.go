package optional

type Optional[T any] struct {
	v *T
}

func (o Optional[T]) IsEmpty() bool {
	return o.v == nil
}

func (o Optional[T]) Get() T {
	if o.IsEmpty() {
		return *new(T)
	}
	return *o.v
}

func Empty[T any]() Optional[T] {
	return Optional[T]{}
}

func Some[T any](v T) Optional[T] {
	return Optional[T]{ v: &v }
}
