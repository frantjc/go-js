package js

// NewComparableArray creates a new ComparableArray from the input array.
func NewComparableArray[T comparable](in []T) ComparableArray[T] {
	out := make(ComparableArray[T], len(in))
	for i, a := range in {
		out[i] = a
	}
	return out
}

// ComparableArray is an array of fmt.Comparable and has methods that
// mimic a subset of the JavaScript array functions.
type ComparableArray[T comparable] []T

// Every tests whether all elements in the array pass the
// test implemented by the provided function. It returns a Boolean value.
func (a ComparableArray[T]) Every(f func(T, int, []T) bool) bool {
	return Every(a, f)
}

// Filter creates a new array with all elements that pass
// the test implemented by the provided function.
func (a ComparableArray[T]) Filter(f func(T, int, []T) bool) ComparableArray[T] {
	return Filter(a, f)
}

// FindIndex returns the index of the first element in the array
// that satisfies the provided testing function. Otherwise, it returns -1,
// indicating that no element passed the test
func (a ComparableArray[T]) FindIndex(f func(T, int, []T) bool) int {
	return FindIndex(a, f)
}

// Find returns the first element in the provided array that satisfies
// the provided testing function. If no values satisfy the testing function,
// the zero value of T is returned.
func (a ComparableArray[T]) Find(f func(T, int, []T) bool) T {
	return Find(a, f)
}

// ForEach executes a provided function once
// for each array element
func (a ComparableArray[T]) ForEach(f func(T, int, []T)) {
	ForEach(a, f)
}

// Map creates a new array populated with the results of
// calling a provided function on every element in the calling array.
func (a ComparableArray[T]) Map(f func(T, int, []T) any) AnyArray[any] {
	return Map(a, f)
}

// TODO solve any typing problem
// ReduceRight applies a function against an accumulator and each value of the array
// (from right-to-left) to reduce it to a single value.
// func (a ComparableArray[T]) ReduceRight(f func(any, T, int, []T) any, initial any) any {
// 	return ReduceRight(a, f, initial)
// }

// TODO solve any typing problem
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
// func (a ComparableArray[T]) Reduce(f func(any, T, int, []T) any, initial any) any {
// 	return Reduce(a, f, initial)
// }

// Reverse creates a new array by reversing the given array.
// The first array element becomes the last, and the last array element becomes the first.
func (a ComparableArray[T]) Reverse() ComparableArray[T] {
	return Reverse(a)
}

// Slice returns a shallow copy of a portion
// of an array into a new array object selected
// from start to end (end not included) where
// start and end represent the index of items in that array.
// The original array will not be modified.
//
// If start<0, it is treated as distance from the end of the array.
// If end<=0, it is treated as distance from the end of the array.
func (a ComparableArray[T]) Slice(start, end int) ComparableArray[T] {
	return Slice(a, start, end)
}

// Some tests whether at least one element in the array passes
// the test implemented by the provided function.
// It returns true if, in the array, it finds an element
// for which the provided function returns true; otherwise
// it returns false. It doesn't modify the array.
func (a ComparableArray[T]) Some(f func(T, int, []T) bool) bool {
	return Some(a, f)
}

// Includes determines whether an array includes a certain value
// among its entries, returning true or false as appropriate.
func (a ComparableArray[T]) Includes(b T) bool {
	return Includes(a, b)
}

// IndexOf returns the first index at which a given element
// can be found in the array, or -1 if it is not present.
func (a ComparableArray[T]) IndexOf(b T) int {
	return IndexOf(a, b)
}

// LastIndexOf returns the last index at which a given element
// can be found in the array, or -1 if it is not present.
// The array is searched backwards, starting at fromIndex.
func (a ComparableArray[T]) LastIndexOf(b T, from int) int {
	return LastIndexOf(a, b, from)
}

// Sort creates a new array by sorting the elements of
// the given array and returns the sorted array.
// The sort order is ascending.
func (a ComparableArray[T]) Sort(f func(a, b T) int) ComparableArray[T] {
	return Sort(a, f)
}

// Unique creates a new array with only unique elements.
func (a ComparableArray[T]) Unique() ComparableArray[T] {
	return Unique(a)
}
