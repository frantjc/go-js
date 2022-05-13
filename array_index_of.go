package js

// IndexOf returns the first index at which a given element
// can be found in the array, or -1 if it is not present.
func IndexOf[T1 comparable](in []T1, a T1) int {
	return FindIndex(in, func(b T1, _ int, _ []T1) bool {
		return a == b
	})
}
