package utils

func IndexOf[T comparable](thing T, arr []T) int {
	for i, t := range arr {
		if t == thing {
			return i
		}
	}
	return -1
}
