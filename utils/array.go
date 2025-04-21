package utils

func Any[T comparable](slice []T, item T) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}

	return false
}

func Map[T any, V any](slice []T, f func(T) V) []V {
	result := make([]V, len(slice))
	for i, v := range slice {
		result[i] = f(v)
	}

	return result
}
