package strs

func Ptr2NumVal(n *int64) int64 {
	if n == nil {
		return 0
	}
	return *n
}
