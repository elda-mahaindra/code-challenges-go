package utils

func Map[T any, U any](slice []T, callbackfn func(T, int, []T) U) []U {
	mapped := []U{}

	for i, v := range slice {
		mapped = append(mapped, callbackfn(v, i, slice))
	}

	return mapped
}
