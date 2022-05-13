package js

// FindIndex returns the index of the first element in the array
// that satisfies the provided testing function. Otherwise, it returns -1,
// indicating that no element passed the test
func FindIndex[T1 any](in []T1, f func(T1, int, []T1) bool) int {
	for i, a := range in {
		if f(a, i, in) {
			return i
		}
	}
	return -1
}
