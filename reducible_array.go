package js

// ReducibleArray is an array of any and has methods that
// mimic the JavaScript array reduce functions.
//
// T1 represents the type of the array, while T2 represents the
// type of the accumulator intended to be used when reducing.
type ReducibleArray[T1, T2 any] AnyArray[T1]

// ReduceRight applies a function against an accumulator and each value of the array
// (from right-to-left) to reduce it to a single value.
func (a ReducibleArray[T1, T2]) ReduceRight(f func(T2, T1, int, []T1) T2, initial T2) T2 {
	return ReduceRight(a, f, initial)
}

// Reduce executes a user-supplied "reducer" callback function
// on each element of the array, in order, passing in the return
// value from the calculation on the preceding element.
// The final result of running the reducer across all
// elements of the array is a single value.
//
// The first time that the callback is run there is no
// "return value of the previous calculation". If supplied,
// an initial value may be used in its place.
// Otherwise the array element at index 0 is used
// as the initial value and iteration starts from
// the next element (index 1 instead of index 0).
func (a ReducibleArray[T1, T2]) Reduce(f func(T2, T1, int, []T1) T2, initial T2) T2 {
	return Reduce(a, f, initial)
}
