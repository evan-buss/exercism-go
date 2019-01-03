package matrix

import "strings"

type Matrix [][]int

// New creates a new matrix object
func New(input string) (Matrix, error) {
	rows := strings.Split(input, "\n")
	col := strings.Split(rows[0], " ")
	var newMatrix [len(rows)][len(col)]Matrix

	for i, line := range rows {
		vals := strings.Split(line, " ")
		for j, val := range vals {

		}

	}

	// row := 0
	// for {
	// 	record, err := r.Read()
	// 	if err == io.EOF {
	// 		break
	// 	}
	// 	if err != nil {
	// 		log.Fatal(err)

	// 	}
	// 	// Convert string array to int array
	// 	newMatrix = append(newMatrix)
	// 	for _, val := range record {
	// 		temp, _ := strconv.Atoi(val)
	// 		newMatrix[row] = append(newMatrix[row], temp)

	// 	}
	// 	row++
	// }

	// fmt.Println(newMatrix)

	return nil, nil
}

func (m Matrix) Rows() [][]int {
	return nil
}

func (m Matrix) Cols() [][]int {
	return nil
}

func (m Matrix) Set(row int, col int, val int) bool {
	return true
}
