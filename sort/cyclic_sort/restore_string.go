package data_structures

func restoreString(s string, indices []int) string {
	bs := []byte(s)
	for i := range indices {
		for indices[i] != i {
			bs[i], bs[indices[i]] = bs[indices[i]], bs[i]
			indices[i], indices[indices[i]] = indices[indices[i]], indices[i]
		}
	}
	return string(bs)
}
