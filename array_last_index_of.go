package js

// LastIndexOf returns the last index at which a given element
// can be found in the array, or -1 if it is not present.
// The array is searched backwards, starting at fromIndex.
func LastIndexOf[T1 comparable](in []T1, a T1) int {
	for i := len(in) - 1; i >= 0; i-- {
		if a == in[i] {
			return i
		}
	}
	return -1
}
