package strain

// Ints is a an int array
type Ints []int

// Lists is an 2 dimensional int array
type Lists [][]int

// Strings is string array
type Strings []string

type strain interface {
	Keep()
	Discard()
}

// Keep returns Ints with only the values that match the conditional
func (ints Ints) Keep(f func(int) bool) Ints {

	var out []int
	for _, val := range ints {
		if f(val) {
			out = append(out, val)
		}
	}
	return out
}

// Discard returns the Ints without values that match the conditional
func (ints Ints) Discard(f func(int) bool) Ints {
	var out []int
	for _, val := range ints {
		if !f(val) {
			out = append(out, val)
		}
	}
	return out
}

// Keep returns the list with values that match the conditional
func (lists Lists) Keep(f func([]int) bool) Lists {

	out := make([][]int, 0, len(lists))
	for _, val := range lists {
		if f(val) {
			out = append(out, val)
		}
	}
	return out
}

// Keep returns the string with values that match the conditional
func (str Strings) Keep(f func(string) bool) Strings {
	out := make([]string, 0, len(str))
	for _, val := range str {
		if f(val) {
			out = append(out, val)
		}
	}
	return out
}
