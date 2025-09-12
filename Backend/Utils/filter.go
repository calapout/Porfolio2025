package Utils

func Filter[T any](items []T, predicate func(T) bool) []T {
	var result []T = make([]T, 0)
	for _, item := range items {
		if predicate(item) {
			result = append(result, item)
		}
	}
	return result
}
