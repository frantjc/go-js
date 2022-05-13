package js

// Filter creates a new array with all elements that pass
// the test implemented by the provided function.
func Filter[T any](in []T, f func(T, int, []T) bool) []T {
	out := []T{}
	for i, a := range in {
		if f(a, i, in) {
			out = append(out, a)
		}
	}
	return out
}
