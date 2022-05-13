package js

// ForEach executes a provided function once
// for each array element
func ForEach[T1 any](in []T1, f func(T1, int, []T1)) {
	for i, a := range in {
		f(a, i, in)
	}
}
