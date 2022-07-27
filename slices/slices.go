package slices

func Contains[E comparable](slice []E, item E) bool {
	for _, e := range slice {
		if e == item {
			return true
		}
	}
	return false
}
