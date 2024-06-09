package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
)

type CacheEntry struct {
	key   string
	value string
}

type Cache struct {
	entries map[string]string
	order   []string
	size    int
}

func NewCache(size int) *Cache {
	return &Cache{
		entries: make(map[string]string),
		order:   make([]string, 0, size),
		size:    size,
	}
}

func (c *Cache) Get(key string) (string, bool) {
	value, found := c.entries[key]
	return value, found
}

func (c *Cache) Put(key, value string) {
	if len(c.entries) == c.size {
		oldestKey := c.order[0]
		c.order = c.order[1:]
		delete(c.entries, oldestKey)
	}
	c.entries[key] = value
	c.order = append(c.order, key)
}

func convertToHex1(binaryString string) (string, error) {
	var hexString string

	for i := len(binaryString); i > 0; i -= 4 {
		startIndex := i - 4
		if startIndex < 0 {
			startIndex = 0
		}
		binChunk := binaryString[startIndex:i]

		decimalValue, err := strconv.ParseInt(binChunk, 2, 64)
		if err != nil {
			return "", fmt.Errorf("invalid binary chunk: %s", err)
		}
		hexString = fmt.Sprintf("%X", decimalValue) + hexString
	}

	return hexString, nil
}

func binarytohexacache(cacheSize int) {
	startTime := time.Now()
	cache := NewCache(cacheSize)

	inputFile, err := os.Open("mat.in")
	if err != nil {
		fmt.Println("Error opening mat.in file:", err)
		return
	}
	defer inputFile.Close()

	outputFile, err := os.Create("mat.in.x")
	if err != nil {
		fmt.Println("Error creating mat.in.x file:", err)
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
		binaryString := parts[1]

		// Check if the binary string is in the cache
		hexString, found := cache.Get(binaryString)
		if !found {
			// If not found, convert and store in cache
			hexString, err = convertToHex1(binaryString)
			if err != nil {
				fmt.Printf("Error converting binary to hex: %v\n", err)
				continue
			}
			cache.Put(binaryString, hexString)
		}

		outputLine := fmt.Sprintf("%s:%s\n", dimensions, hexString)
		_, err = io.WriteString(outputFile, outputLine)
		if err != nil {
			fmt.Println("Error writing to file:", err)
			return
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading mat.in file:", err)
		return
	}

	fmt.Println("Conversion completed successfully. Output written to mat.in.x")

	elapsedTime := time.Since(startTime) // Calculate elapsed time
	fmt.Printf("Program execution time: %s\n", elapsedTime)
}
