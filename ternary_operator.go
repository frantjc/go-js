package js

// Ternary returns the first T1 if the conditional evaluates to true,
// otherwise it returns the second T1
func Ternary[T1 any](condition bool, a, b T1) T1 {
	if condition {
		return a
	}
	return b
}
