package utils

func IndexOf[T comparable](thing T, arr []T) int {
	for i, t := range arr {
		if t == thing {
			return i
		}
	}
	return -1
}

func Map[T, U any](arr []T, f func(T) U) []U {
	var result []U
	for _, a := range arr {
		result = append(result, f(a))
	}
	return result
}

func Find[T any](f func(T) bool, arr []T) *T {
	for _, a := range arr {
		if f(a) {
			return &a
		}
	}
	return nil
}
