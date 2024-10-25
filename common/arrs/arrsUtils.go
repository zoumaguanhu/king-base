package arrs

func IsDefault[T any](a []T) bool {
	return len(a) <= 0
}
func IsNotDefault[T any](a []T) bool {
	return !IsDefault(a)
}
func PtrIsNotNull[T any](a *[]T) bool {
	return a != nil || IsNotDefault(*a)
}
func PtrIsPtrNotNull[T any](a *[]*T) bool {
	return a != nil || len(*a) > 0
}
func PtrIsNull[T any](a *[]T) bool {
	return a == nil || IsDefault(*a)
}
func PtrIsPtrNull[T any](a *[]*T) bool {
	return a == nil || len(*a) <= 0
}
