package helper

// PtrToInt returns a pointer to a copy of the given integer
func PtrToInt(i int) *int {
	r := new(int)
	*r = i
	return r
}
