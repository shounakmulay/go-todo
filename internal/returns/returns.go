package returns

func ValueOrError[T any](value T, err error) (T, error) {
	if err != nil {
		return *new(T), err
	}
	return value, nil
}

func ErrorOrElse[T any](err error, value func() T) (T, error) {
	if err != nil {
		return *new(T), err
	}
	return value(), nil
}
