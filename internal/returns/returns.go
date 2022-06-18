package returns

func ErrorOrValue[T any](err error, value T) (T, error) {
	if err != nil {
		return *new(T), err
	}
	return value, nil
}
