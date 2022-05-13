package js

// ComparableArray is an array of fmt.Comparable and has methods that
// mimic a subset of the JavaScript array functions that are
// unique to comparables.
type ComparableArray[T comparable] AnyArray[T]

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

// Unique creates a new array with only unique elements from the given array.
func (a ComparableArray[T]) Unique() ComparableArray[T] {
	return Unique(a)
}
