package js

// Some tests whether at least one element in the array passes
// the test implemented by the provided function.
// It returns true if, in the array, it finds an element
// for which the provided function returns true; otherwise
// it returns false. It doesn't modify the array.
func Some[T1 any](in []T1, f func(T1, int, []T1) bool) bool {
	for i, a := range in {
		if f(a, i, in) {
			return true
		}
	}
	return false
}
