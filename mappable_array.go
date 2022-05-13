package js

// MappableArray is an array of any and has methods that
// mimic the JavaScript array map function.
//
// T1 represents the type of the array, while T2 represents the
// type of the array intended to be mapped to
type MappableArray[T1, T2 any] AnyArray[T1]

// Map creates a new array populated with the results of
// calling a provided function on every element in the calling array.
func (a MappableArray[T1, T2]) Map(f func(T1, int, []T1) T2) AnyArray[T2] {
	return Map(a, f)
}
