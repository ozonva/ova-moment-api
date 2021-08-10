package slice

import (
	"github.com/ozonva/ova-moment-api/internal/utils"
)

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
