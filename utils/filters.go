package utils

func Filter[T comparable](slice []T, filter func(T) bool) []T {
	filtered := []T{}

	for _, v := range slice {
		if filter(v) {
			filtered = append(filtered, v)
		}
	}

	return filtered
}
