package strs

func IsEmpty(s *string) bool {
	return s == nil || *s == ""
}
func NotEmpty(s *string) bool {
	return !IsEmpty(s)
}
