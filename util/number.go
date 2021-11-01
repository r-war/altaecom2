package utils

func NullIntScan(a *int) int {
	if a != nil {
		return *a
	}

	return 0
}
