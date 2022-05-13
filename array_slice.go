package js

// Slice returns a shallow copy of a portion
// of an array into a new array object selected
// from start to end (end not included) where
// start and end represent the index of items in that array.
// The original array will not be modified.
//
// If start<0, it is treated as distance from the end of the array
// If end<=0, it is treated as distance from the end of the array
func Slice[T1 any](in []T1, start, end int) []T1 {
	if start > end {
		return make([]T1, 0)
	}
	if start < 0 {
		start = len(in) + start
	}
	if end <= 0 {
		end = len(in) + end
	}
	out := make([]T1, end-start)
	for k, a := range in[start:end] {
		out[k] = a
	}
	return out
}
