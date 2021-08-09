package slice

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestChunksInt(t *testing.T) {
	type successCase struct {
		slice     []int
		chunkSize int
		chunks    [][]int
	}

	type panicCase struct {
		slice     []int
		chunkSize int
		error     string
	}

	var successCases = []successCase{
		{
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14},
			4,
			[][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14}},
		},
		{
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14},
			2,
			[][]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}, {9, 10}, {11, 12}, {13, 14}},
		},
		{
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14},
			7,
			[][]int{{1, 2, 3, 4, 5, 6, 7}, {8, 9, 10, 11, 12, 13, 14}},
		},
	}

	var panicCases = []panicCase{
		{[]int{1, 2, 3, 4, 5}, 0, "Chunk size must be a positive integer"},
		{[]int{}, -2, "Chunk size must be a positive integer"},
	}

	for _, successCase := range successCases {
		res := ChunksInt(successCase.slice, successCase.chunkSize)
		assert.Equal(t, successCase.chunks, res, "The two slices must be same")
	}

	for _, panicCase := range panicCases {
		assert.PanicsWithValue(t, panicCase.error, func() {
			ChunksInt(panicCase.slice, panicCase.chunkSize)
		}, "The two errors must be same")
	}
}

func TestChunksStr(t *testing.T) {
	type successCase struct {
		slice     []string
		chunkSize int
		chunks    [][]string
	}

	type panicCase struct {
		slice     []string
		chunkSize int
		error     string
	}

	var successCases = []successCase{
		{
			[]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n"},
			4,
			[][]string{{"a", "b", "c", "d"}, {"e", "f", "g", "h"}, {"i", "j", "k", "l"}, {"m", "n"}},
		},
		{
			[]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n"},
			2,
			[][]string{{"a", "b"}, {"c", "d"}, {"e", "f"}, {"g", "h"}, {"i", "j"}, {"k", "l"}, {"m", "n"}},
		},
		{
			[]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n"},
			7,
			[][]string{{"a", "b", "c", "d", "e", "f", "g"}, {"h", "i", "j", "k", "l", "m", "n"}},
		},
	}

	var panicCases = []panicCase{
		{[]string{"a", "b", "c", "d", "e"}, 0, "Chunk size must be a positive integer"},
		{[]string{}, -2, "Chunk size must be a positive integer"},
	}

	for _, successCase := range successCases {
		res := ChunksStr(successCase.slice, successCase.chunkSize)
		assert.Equal(t, successCase.chunks, res, "The two slices must be same")
	}

	for _, panicCase := range panicCases {
		assert.PanicsWithValue(t, panicCase.error, func() {
			ChunksStr(panicCase.slice, panicCase.chunkSize)
		}, "The two errors must be same")
	}
}
