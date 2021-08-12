package _map

// FlipStrStr exchanges all keys with their associated values in the map of string-string
//
func FlipStrStr(m map[string]string) map[string]string {
	flippedMap := make(map[string]string)
	for key, v := range m {
		if _, found := flippedMap[v]; found {
			panic("Values of input map are not unique")
		}

		flippedMap[v] = key
	}

	return flippedMap
}

// FlipIntInt exchanges all keys with their associated values in the map of int-int
//
func FlipIntInt(m map[int]int) map[int]int {
	flippedMap := make(map[int]int)
	for key, v := range m {
		if _, found := flippedMap[v]; found {
			panic("Values of input map are not unique")
		}

		flippedMap[v] = key
	}

	return flippedMap
}

// FlipStrInt exchanges all keys with their associated values in the map of string-int
//
func FlipStrInt(m map[string]int) map[int]string {
	flippedMap := make(map[int]string)
	for key, v := range m {
		if _, found := flippedMap[v]; found {
			panic("Values of input map are not unique")
		}

		flippedMap[v] = key
	}

	return flippedMap
}

// FlipIntStr exchanges all keys with their associated values in the map of int-string
//
func FlipIntStr(m map[int]string) map[string]int {
	flippedMap := make(map[string]int)
	for key, v := range m {
		if _, found := flippedMap[v]; found {
			panic("Values of input map are not unique")
		}

		flippedMap[v] = key
	}

	return flippedMap
}
