package _map

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFlipStrStr(t *testing.T) {
	type successCase struct {
		m          map[string]string
		flippedMap map[string]string
	}

	type panicCase struct {
		m     map[string]string
		error string
	}

	successCases := []successCase{
		{map[string]string{"a": "z", "b": "y", "c": "x"}, map[string]string{"z": "a", "y": "b", "x": "c"}},
		{map[string]string{}, map[string]string{}},
	}

	panicCases := []panicCase{
		{map[string]string{"a": "z", "b": "y", "c": "z"}, "Values of input map are not unique"},
	}

	for _, successCase := range successCases {
		assert.Equal(t, successCase.flippedMap, FlipStrStr(successCase.m), "The two maps must be same")
	}

	for _, panicCase := range panicCases {
		assert.PanicsWithValue(t, panicCase.error, func() {
			FlipStrStr(panicCase.m)
		}, "The two errors must be same")
	}
}

func TestFlipIntInt(t *testing.T) {
	type successCase struct {
		m          map[int]int
		flippedMap map[int]int
	}

	type panicCase struct {
		m     map[int]int
		error string
	}

	successCases := []successCase{
		{map[int]int{0: 9, 1: 8, 2: 7}, map[int]int{9: 0, 8: 1, 7: 2}},
		{map[int]int{}, map[int]int{}},
	}

	panicCases := []panicCase{
		{map[int]int{0: 9, 1: 8, 2: 9}, "Values of input map are not unique"},
	}

	for _, successCase := range successCases {
		assert.Equal(t, successCase.flippedMap, FlipIntInt(successCase.m), "The two maps must be same")
	}

	for _, panicCase := range panicCases {
		assert.PanicsWithValue(t, panicCase.error, func() {
			FlipIntInt(panicCase.m)
		}, "The two errors must be same")
	}
}

func TestFlipStrInt(t *testing.T) {
	type successCase struct {
		m          map[string]int
		flippedMap map[int]string
	}

	type panicCase struct {
		m     map[string]int
		error string
	}

	successCases := []successCase{
		{map[string]int{"a": 9, "b": 8, "c": 7}, map[int]string{9: "a", 8: "b", 7: "c"}},
		{map[string]int{}, map[int]string{}},
	}

	panicCases := []panicCase{
		{map[string]int{"a": 9, "b": 8, "c": 9}, "Values of input map are not unique"},
	}

	for _, successCase := range successCases {
		assert.Equal(t, successCase.flippedMap, FlipStrInt(successCase.m), "The two maps must be same")
	}

	for _, panicCase := range panicCases {
		assert.PanicsWithValue(t, panicCase.error, func() {
			FlipStrInt(panicCase.m)
		}, "The two errors must be same")
	}
}

func TestFlipIntStr(t *testing.T) {
	type successCase struct {
		m          map[int]string
		flippedMap map[string]int
	}

	type panicCase struct {
		m     map[int]string
		error string
	}

	successCases := []successCase{
		{map[int]string{0: "z", 1: "y", 2: "x"}, map[string]int{"z": 0, "y": 1, "x": 2}},
		{map[int]string{}, map[string]int{}},
	}

	panicCases := []panicCase{
		{map[int]string{0: "z", 1: "y", 2: "z"}, "Values of input map are not unique"},
	}

	for _, successCase := range successCases {
		assert.Equal(t, successCase.flippedMap, FlipIntStr(successCase.m), "The two maps must be same")
	}

	for _, panicCase := range panicCases {
		assert.PanicsWithValue(t, panicCase.error, func() {
			FlipIntStr(panicCase.m)
		}, "The two errors must be same")
	}
}
