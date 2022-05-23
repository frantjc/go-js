package js

// Coalesce is meant to represent the '||' operator e.g.
//
//  const myVar = otherVar || 'default';
//
// It returns the first non-zero-valued instance of T that it finds,
// or the zero value of T if it finds none.
func Coalesce[T comparable](ts ...T) T {
	for _, t := range ts {
		if t != *new(T) {
			return t
		}
	}
	return *new(T)
}
