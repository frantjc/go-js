package js

import "fmt"

// NewStringerArray creates a new StringerArray from the input array.
func NewStringerArray[T fmt.Stringer](in []T) StringerArray[T] {
	out := make(StringerArray[T], len(in))
	for i, a := range in {
		out[i] = a
	}
	return out
}

// StringerArray is an array of fmt.Stringer and has methods that
// mimic a subset of the JavaScript array functions.
type StringerArray[T fmt.Stringer] []T

// Every tests whether all elements in the array pass the
// test implemented by the provided function. It returns a Boolean value.
func (a StringerArray[T]) Every(f func(T, int, []T) bool) bool {
	return Every(a, f)
}

// Filter creates a new array with all elements that pass
// the test implemented by the provided function.
func (a StringerArray[T]) Filter(f func(T, int, []T) bool) StringerArray[T] {
	return Filter(a, f)
}

// FindIndex returns the index of the first element in the array
// that satisfies the provided testing function. Otherwise, it returns -1,
// indicating that no element passed the test
func (a StringerArray[T]) FindIndex(f func(T, int, []T) bool) int {
	return FindIndex(a, f)
}

// Find returns the first element in the provided array that satisfies
// the provided testing function. If no values satisfy the testing function,
// the zero value of T is returned.
func (a StringerArray[T]) Find(f func(T, int, []T) bool) T {
	return Find(a, f)
}

// ForEach executes a provided function once
// for each array element
func (a StringerArray[T]) ForEach(f func(T, int, []T)) {
	ForEach(a, f)
}

// Map creates a new array populated with the results of
// calling a provided function on every element in the calling array.
func (a StringerArray[T]) Map(f func(T, int, []T) any) AnyArray[any] {
	return Map(a, f)
}

// ReduceRight applies a function against an accumulator and each value of the array
// (from right-to-left) to reduce it to a single value.
func (a StringerArray[T]) ReduceRight(f func(any, T, int, []T) any, initial any) any {
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
func (a StringerArray[T]) Reduce(f func(any, T, int, []T) any, initial any) any {
	return Reduce(a, f, initial)
}

// Reverse creates a new array by reversing the given array.
// The first array element becomes the last, and the last array element becomes the first.
func (a StringerArray[T]) Reverse() StringerArray[T] {
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
func (a StringerArray[T]) Slice(start, end int) StringerArray[T] {
	return Slice(a, start, end)
}

// Some tests whether at least one element in the array passes
// the test implemented by the provided function.
// It returns true if, in the array, it finds an element
// for which the provided function returns true; otherwise
// it returns false. It doesn't modify the array.
func (a StringerArray[T]) Some(f func(T, int, []T) bool) bool {
	return Some(a, f)
}

// Join creates and returns a new string by concatenating all of the elements
// in an array (or an array-like object), separated by commas or a specified
// separator string. If the array has only one item, then that item will be
// returned without using the separator.
func (a StringerArray[T]) Join(separator string) string {
	return Join(a, separator)
}
