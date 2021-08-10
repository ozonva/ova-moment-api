package _map

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
