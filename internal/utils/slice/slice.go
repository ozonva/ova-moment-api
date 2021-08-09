package slice

func ChunksInt(slice []int, chunkSize int) [][]int {
	if chunkSize <= 0 {
		panic("Chunk size must be a positive integer")
	}

	chunks := make([][]int, (len(slice)+chunkSize-1)/chunkSize)
	for i := 0; i < len(chunks); i++ {
		chunks[i] = slice[i*chunkSize : minInt(i*chunkSize+chunkSize, len(slice))]
	}

	return chunks
}

func ChunksStr(slice []string, chunkSize int) [][]string {
	if chunkSize <= 0 {
		panic("Chunk size must be a positive integer")
	}

	chunks := make([][]string, (len(slice)+chunkSize-1)/chunkSize)
	for i := 0; i < len(chunks); i++ {
		chunks[i] = slice[i*chunkSize : minInt(i*chunkSize+chunkSize, len(slice))]
	}

	return chunks
}

func minInt(a int, b int) int {
	if a <= b {
		return a
	}
	return b
}
