package main

func Map[T interface{}, U interface{}](list []T, fn func(T) U) []U {
	out := make([]U, len(list))

	for i, v := range list {
		out[i] = fn(v)
	}

	return out
}
