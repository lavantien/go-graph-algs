package common

func Reverse[T any](list []T) {
	for i, j := 0, len(list)-1; i < j; i, j = i+1, j-1 {
		list[i], list[j] = list[j], list[i]
	}
}

func Max[T int](a T, b T) T {
	if a > b {
		return a
	}
	return b
}

func Min[T int](a T, b T) T {
	if a < b {
		return a
	}
	return b
}
