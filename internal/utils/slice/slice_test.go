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

	successCases := []successCase{
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

	panicCases := []panicCase{
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

	successCases := []successCase{
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

	panicCases := []panicCase{
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

func TestDiffInt(t *testing.T) {
	type successCase struct {
		minuend    []int
		subtrahend []int
		difference []int
	}

	successCases := []successCase{
		{
			[]int{5, 1, 2, 8, 0, 1, 1, 3, 9, 0, 12},
			[]int{1, 13, 6, 5, 5},
			[]int{2, 8, 0, 3, 9, 0, 12},
		},
		{
			[]int{5, 1, 2, 8, 0, 1, 1, 3, 9, 0, 12},
			[]int{1, 13, 6, 5, 5, 8, 2, 0, 12, 9, 3, 9},
			[]int{},
		},
		{
			[]int{5, 1, 2, 8, 0, 1, 1, 3, 9, 0, 12},
			[]int{24, 25},
			[]int{5, 1, 2, 8, 0, 1, 1, 3, 9, 0, 12},
		},
		{
			[]int{5, 1, 2, 8, 0, 1, 1, 3, 9, 0, 12},
			[]int{},
			[]int{5, 1, 2, 8, 0, 1, 1, 3, 9, 0, 12},
		},
	}

	for _, successCase := range successCases {
		res := DiffInt(successCase.minuend, successCase.subtrahend...)
		assert.Equal(t, successCase.difference, res, "The two slices must be same")
	}
}

func TestDiffStr(t *testing.T) {
	type successCase struct {
		minuend    []string
		subtrahend []string
		difference []string
	}

	successCases := []successCase{
		{
			[]string{"f", "b", "c", "i", "a", "b", "b", "d", "j", "a", "m"},
			[]string{"b", "n", "g", "f", "f"},
			[]string{"c", "i", "a", "d", "j", "a", "m"},
		},
		{
			[]string{"f", "b", "c", "i", "a", "b", "b", "d", "j", "a", "m"},
			[]string{"b", "n", "g", "f", "f", "i", "c", "a", "m", "j", "d", "j"},
			[]string{},
		},
		{
			[]string{"f", "b", "c", "i", "a", "b", "b", "d", "j", "a", "m"},
			[]string{"y", "z"},
			[]string{"f", "b", "c", "i", "a", "b", "b", "d", "j", "a", "m"},
		},
		{
			[]string{"f", "b", "c", "i", "a", "b", "b", "d", "j", "a", "m"},
			[]string{},
			[]string{"f", "b", "c", "i", "a", "b", "b", "d", "j", "a", "m"},
		},
	}

	for _, successCase := range successCases {
		res := DiffStr(successCase.minuend, successCase.subtrahend...)
		assert.Equal(t, successCase.difference, res, "The two slices must be same")
	}
}
