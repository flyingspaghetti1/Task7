package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func convertToBinary(hexString string) (string, error) {
	var binaryString string

	for _, r := range hexString {
		decimalValue, err := strconv.ParseInt(string(r), 16, 64)
		if err != nil {
			return "", fmt.Errorf("invalid hex character: %s", err)
		}
		binaryChunk := fmt.Sprintf("%04b", decimalValue)
		binaryString += binaryChunk
	}

	return binaryString, nil
}

func hexatobinary() {
	inputFile, err := os.Open("mat.in.x")
	if err != nil {
		fmt.Println("Error opening mat.in.x file:", err)
		return
	}
	defer inputFile.Close()

	outputFile, err := os.Create("mat.in")
	if err != nil {
		fmt.Println("Error creating mat.in file:", err)
		return
	}
	defer outputFile.Close()

	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			continue
		}
		parts := strings.Split(line, ":")
		if len(parts) != 2 {
			fmt.Printf("Skipping invalid line: %s\n", line)
			continue
		}
		dimensions := parts[0]
		hexString := parts[1]

		binaryString, err := convertToBinary(hexString)
		if err != nil {
			fmt.Printf("Error converting hex to binary: %v\n", err)
			continue
		}

		outputLine := fmt.Sprintf("%s:%s\n", dimensions, binaryString)
		_, err = io.WriteString(outputFile, outputLine)
		if err != nil {
			fmt.Println("Error writing to file:", err)
			return
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading mat.in.x file:", err)
		return
	}

	fmt.Println("Conversion completed successfully. Output written to mat.in")
}
