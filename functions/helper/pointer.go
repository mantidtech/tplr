package helper

// PtrTo returns a pointer to a copy of the given parameter
func PtrTo[T any](t T) *T {
	r := new(T)
	*r = t
	return r
}
