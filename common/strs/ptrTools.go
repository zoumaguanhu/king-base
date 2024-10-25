package strs

func Ptr2NumVal(n *int64) int64 {
	if n == nil {
		return 0
	}
	return *n
}
func ArrPtrIsNull[T any](p []*T) bool {
	return p == nil || len(p) < 0
}
func PtrArrPtrIsNull[T any](p *[]*T) bool {
	return p == nil || len(*p) < 0
}
func PtrArrIsNull[T any](p *[]T) bool {
	return p == nil || len(*p) < 0
}
