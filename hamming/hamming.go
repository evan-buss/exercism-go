package hamming

import "errors"

// Distance returns the hamming distance between two DNA strings
func Distance(a, b string) (int, error) {

	if len(a) != len(b) {
		return 0, errors.New("unequal lengths")
	}
	difCount := 0
	for i := range a {
		if a[i] != b[i] {
			difCount++
		}
	}
	return difCount, nil
}
