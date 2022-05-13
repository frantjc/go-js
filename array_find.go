package js

// Find returns the first element in the provided array that satisfies
// the provided testing function. If no values satisfy the testing function,
// the zero value of T1 is returned.
func Find[T1 any](in []T1, f func(T1, int, []T1) bool) T1 {
	for i, a := range in {
		if f(a, i, in) {
			return a
		}
	}
	return *new(T1)
}
