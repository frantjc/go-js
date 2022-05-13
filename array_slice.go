package js

// Slice returns a shallow copy of a portion
// of an array into a new array object selected
// from start to end (end not included) where
// start and end represent the index of items in that array.
// The original array will not be modified.
//
// If end <=0, it is treated as distance from the end of the array
func Slice[T1 any](in []T1, start, end int) []T1 {
	j := end
	if j <= 0 {
		j = len(in) + j
	}
	out := make([]T1, j-start)
	for k, a := range in[start:j] {
		out[k] = a
	}
	return out
}
