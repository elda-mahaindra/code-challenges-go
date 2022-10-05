package utils

func Reduce[T comparable, U any](slice []T, initial U, callbackfn func(U, T, int, []T) U) U {
	reduced := initial

	for i, v := range slice {
		reduced = callbackfn(reduced, v, i, slice)
	}

	return reduced
}
