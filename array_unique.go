package js

// Unique creates a new array with only unique elements
func Unique[T1 comparable](in []T1) []T1 {
	return Filter(in, func(a T1, i int, self []T1) bool {
		return i == IndexOf(self, a)
	})
}
