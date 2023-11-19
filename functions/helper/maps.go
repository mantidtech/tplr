package helper

// Combine a slice of maps into a single map, with key collisions, later overwrites earlier
func Combine[M ~map[K]V, K comparable, V any](maps ...M) M {
	res := make(M)
	for _, m := range maps {
		for k, v := range m {
			res[k] = v
		}
	}
	return res
}
