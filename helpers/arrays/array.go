package arrays

// AppendIfMissing appenda i if not found in slice.
func AppendIfMissing(slice []string, i string) []string {
	for _, ele := range slice {
		if ele == i {
			return slice
		}
	}
	return append(slice, i)
}

// AppendAllIfMissing appenda i if not found in slice.
func AppendAllIfMissing(slice []string, inputs ...string) []string {
	for _, s := range inputs {
		slice = AppendIfMissing(slice, s)
	}
	return slice
}

// RemoveFromArray removes a string object from the given slice
func RemoveFromArray(slice []string, input string) []string {
	var output []string
	for i, item := range slice {
		if item == input {
			output = append(slice[:i], slice[i+1:]...)
			break
		}
	}
	return output
}
