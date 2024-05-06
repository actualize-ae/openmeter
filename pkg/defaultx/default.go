package defaultx

func WithDefault[T any](value *T, def T) T {
	if value != nil {
		return *value
	}

	return def
}

func ToPointerOrNil[T comparable](v T) *T {
	var defaultValue T

	if v == defaultValue {
		return nil
	}

	return &v
}
