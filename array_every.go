package js

// Every tests whether all elements in the array pass the
// test implemented by the provided function. It returns a Boolean value.
func Every[T any](in []T, f func(T, int, []T) bool) bool {
	for i, a := range in {
		if !f(a, i, in) {
			return false
		}
	}
	return true
}
