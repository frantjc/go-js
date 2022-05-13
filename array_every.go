package js

// Every tests whether all elements in the array pass the
// test implemented by the provided function. It returns a Boolean value.
func Every[T1 any](in []T1, f func(T1, int, []T1) bool) bool {
	for i, a := range in {
		if !f(a, i, in) {
			return false
		}
	}
	return true
}
