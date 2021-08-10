package slice

import (
	"github.com/ozonva/ova-moment-api/internal/utils"
)

// ChunksInt Splits the slice of ints into chunks
//
func ChunksInt(slice []int, chunkSize int) [][]int {
	if chunkSize <= 0 {
		panic("Chunk size must be a positive integer")
	}

	chunks := make([][]int, (len(slice)+chunkSize-1)/chunkSize)
	for i := 0; i < len(chunks); i++ {
		chunks[i] = slice[i*chunkSize : utils.MinInt(i*chunkSize+chunkSize, len(slice))]
	}

	return chunks
}

// ChunksStr Splits the slice of strings into chunks
//
func ChunksStr(slice []string, chunkSize int) [][]string {
	if chunkSize <= 0 {
		panic("Chunk size must be a positive integer")
	}

	chunks := make([][]string, (len(slice)+chunkSize-1)/chunkSize)
	for i := 0; i < len(chunks); i++ {
		chunks[i] = slice[i*chunkSize : utils.MinInt(i*chunkSize+chunkSize, len(slice))]
	}

	return chunks
}

// DiffInt Computes the difference of slice of ints and elements.
//
func DiffInt(minuend []int, subtrahend ...int) []int {
	var stub struct{}
	subtrahendMap := make(map[int]struct{}, len(subtrahend))
	for _, v := range subtrahend {
		subtrahendMap[v] = stub
	}

	difference := make([]int, 0)
	for _, v := range minuend {
		if _, found := subtrahendMap[v]; !found {
			difference = append(difference, v)
		}
	}

	return difference
}

// DiffStr Computes the difference of slice of strings and elements.
//
func DiffStr(minuend []string, subtrahend ...string) []string {
	var stub struct{}
	subtrahendMap := make(map[string]struct{}, len(subtrahend))
	for _, v := range subtrahend {
		subtrahendMap[v] = stub
	}

	difference := make([]string, 0)
	for _, v := range minuend {
		if _, found := subtrahendMap[v]; !found {
			difference = append(difference, v)
		}
	}

	return difference
}
