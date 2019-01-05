package matrix

import (
	"encoding/csv"
	"errors"
	"strconv"
	"strings"
)

// Matrix represents a matrix of integers
type Matrix [][]int

// New creates a new matrix object
func New(input string) (Matrix, error) {

	// Make sure there are no empty rows
	if input[0] == '\n' || input[len(input)-1] == '\n' {
		return nil, errors.New("Empty first or last row")
	} else if strings.Contains(input, "\n\n") {
		return nil, errors.New("Empty inside row")
	}

	// Create new csv reader
	r := csv.NewReader(strings.NewReader(input))
	r.Comma = ' '
	r.Comment = '#'
	r.TrimLeadingSpace = true
	r.FieldsPerRecord = 0 //Gets the number of fields from the first line

	// Read entire into [][]string
	arr, err := r.ReadAll()
	if err != nil {
		return nil, errors.New("Could not parse string")
	}

	// Create new matrix the same size as the input
	m := make(Matrix, len(arr))
	for i := range m {
		m[i] = make([]int, len(arr[0]))
	}

	// Fill the new matrix with int conversions of the original
	for i, row := range arr {
		for j, val := range row {
			m[i][j], err = strconv.Atoi(val)
			if err != nil {
				return nil, errors.New("value too large")
			}

		}
	}
	return m, nil
}

// Rows returns the rows of the given matrix
func (m Matrix) Rows() [][]int {
	// Create new 2D array the same size as the input
	temp := make([][]int, len(m))
	for i := range temp {
		temp[i] = make([]int, len(m[0]))
	}
	// Transfer the values to the copy
	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[0]); j++ {
			temp[i][j] = m[i][j]
		}
	}

	return temp
}

// Cols returns the columns of the given matrix
func (m Matrix) Cols() [][]int {
	// Create new 2D array the same size as the input
	temp := make([][]int, len(m[0]))
	for i := range temp {
		temp[i] = make([]int, len(m))
	}
	// Swap the items
	for i, s := range m {
		for j, e := range s {
			temp[j][i] = e
		}
	}
	return temp
}

// Set allows you to set a specific element of the matrix
func (m *Matrix) Set(row int, col int, val int) bool {
	// Make sure that the index values are within the matrix's range
	if row >= 0 && row < len(*m) && col >= 0 && col < len((*m)[0]) {
		(*m)[row][col] = val
		return true
	}
	return false
}
