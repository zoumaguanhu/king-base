package arrs

func IsDefault[T any](a []T) bool {
	return len(a) <= 0
}
func IsNotDefault[T any](a []T) bool {
	return !IsDefault(a)
}
func IsArrNotNull[T any](a []*T) bool {
	return !IsArrNull(a)
}
func IsArrNull[T any](a []*T) bool {
	return len(a) <= 0
}
func PtrArrIsNotNull[T any](a *[]T) bool {
	return a != nil || IsNotDefault(*a)
}
func PtrArrIsPtrNotNull[T any](a *[]*T) bool {
	return a != nil || len(*a) > 0
}
func PtrArrIsNull[T any](a *[]T) bool {
	return a == nil || IsDefault(*a)
}
func PtrArrIsPtrNull[T any](a *[]*T) bool {
	return a == nil || len(*a) <= 0
}
