package js

// Includes determines whether an array includes a certain value
// among its entries, returning true or false as appropriate.
func Includes[T1 comparable](in []T1, a T1) bool {
	for _, b := range in {
		if a == b {
			return true
		}
	}
	return false
}
