package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func generateMatrix(rows, columns int) [][]int {
	matrix := make([][]int, rows)
	for i := range matrix {
		matrix[i] = make([]int, columns)
		for j := range matrix[i] {
			matrix[i][j] = rand.Intn(2)
		}
	}
	return matrix
}

func matrixToString(matrix [][]int) string {
	var sb strings.Builder
	for _, row := range matrix {
		for _, cell := range row {
			sb.WriteString(strconv.Itoa(cell))
		}
	}
	return sb.String()
}

func writeToFIle(filename string, matrices []struct {
	rows, columns int
	matrix        [][]int
}) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	for _, m := range matrices {
		matrixStr := matrixToString(m.matrix)
		_, err = file.WriteString(fmt.Sprintf("%dx%d:%s\n", m.rows, m.columns, matrixStr))
		if err != nil {
			return err
		}
	}
	return nil
}

func duplicates() {
	numMatrices := 300000
	numDuplicates := 250000

	uniqueMatrices := make(map[string]struct{ rows, columns int })
	var matrices []struct {
		rows, columns int
		matrix        [][]int
	}

	for len(uniqueMatrices) < numMatrices-numDuplicates {
		rows := rand.Intn(6) + 5
		columns := rand.Intn(6) + 5
		matrix := generateMatrix(rows, columns)
		matrixStr := matrixToString(matrix)

		if _, ok := uniqueMatrices[matrixStr]; !ok {
			uniqueMatrices[matrixStr] = struct{ rows, columns int }{rows, columns}
			matrices = append(matrices, struct {
				rows, columns int
				matrix        [][]int
			}{rows, columns, matrix})
		}
	}

	for i := 0; i < numDuplicates; i++ {
		var randomMatrix struct {
			rows, columns int
			matrix        [][]int
		}
		for matrixStr, m := range uniqueMatrices {
			randomMatrix = struct {
				rows, columns int
				matrix        [][]int
			}{m.rows, m.columns, generateMatrix(m.rows, m.columns)}
			randomMatrix.matrix = generateMatrix(m.rows, m.columns)
			for i := 0; i < m.rows; i++ {
				for j := 0; j < m.columns; j++ {
					randomMatrix.matrix[i][j], _ = strconv.Atoi(string(matrixStr[i*m.columns+j]))
				}
			}
			break
		}
		matrices = append(matrices, randomMatrix)
	}

	rand.Shuffle(len(matrices), func(i, j int) {
		matrices[i], matrices[j] = matrices[j], matrices[i]
	})

	err := writeToFIle("mat.in", matrices)
	if err != nil {
		fmt.Println("Error writing to file:", err)
	}
}
