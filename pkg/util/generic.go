package util

// DefaultIfEmpty returns defaultValue if value is the zero value for its type.
// Otherwise, it returns value.
// This function works for any comparable type.
// Example usage:
//
//	result := DefaultIfEmpty(someString, "default")
//	result := DefaultIfEmpty(someInt, 42)
func DefaultIfEmpty[T comparable](value, defaultValue T) T {
	var zero T
	if value == zero {
		return defaultValue
	}
	return value
}
