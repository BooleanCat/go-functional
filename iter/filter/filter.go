// This package contains functions intended for use with [iter.Filter].
package filter

// IsEven returns true when the provided integer is even.
func IsEven[T ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr | ~int | ~int8 | ~int16 | ~int32 | ~int64](integer T) bool {
	return integer%2 == 0
}

// IsOdd returns true when the provided integer is odd.
func IsOdd[T ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr | ~int | ~int8 | ~int16 | ~int32 | ~int64](integer T) bool {
	return integer%2 != 0
}
