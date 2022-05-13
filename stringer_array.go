package js

import "fmt"

// StringerArray is an array of fmt.Stringer and has methods that
// mimic a subset of the JavaScript array functions that are
// unique to fmt.Stringers.
type StringerArray[T fmt.Stringer] AnyArray[T]

// Join creates and returns a new string by concatenating all of the elements
// in an array (or an array-like object), separated by commas or a specified
// separator string. If the array has only one item, then that item will be
// returned without using the separator.
func (a StringerArray[T]) Join(separator string) string {
	return Join(a, separator)
}
