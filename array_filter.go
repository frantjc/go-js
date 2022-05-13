package js

// Filter creates a new array with all elements that pass
// the test implemented by the provided function.
func Filter[T1 any](in []T1, f func(T1, int, []T1) bool) []T1 {
	out := []T1{}
	for i, a := range in {
		if f(a, i, in) {
			out = append(out, a)
		}
	}
	return out
}
