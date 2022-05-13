package js

// Sort creates a new array by sorting the elements of
// the given array and returns the sorted array.
// The sort order is ascending.
func Sort[T1 comparable](in []T1, f func(a, b T1) int) []T1 {
	out := in
	k := len(out)
	for i := 0; i < k; i++ {
		for j := 0; j < k-1; j++ {
			if f(out[j], out[j+1]) > 0 {
				out[j], out[j+1] = out[j+1], out[j]
			}
		}
	}
	return out
}
