package day47

func LongestSubstringWithoutDuplication(str string) int {
	longest := 0
	runeStr := []rune(str)
	position := make([]rune, 26)
	for start := 0; start < len(runeStr); start++ {
		for end := start; end < len(runeStr); end++ {
			count := end - start + 1
			substring := runeStr[start : count+start]
			if !hasDuplication(substring, position) {
				if count > longest {
					longest = count
				}
			} else {
				break
			}
		}

	}
	return longest
}

func hasDuplication(substring []rune, position []rune) bool {
	for i := 0; i < 26; i++ {
		position[i] = -1
	}

	for i := 0; i < len(substring); i++ {
		indexInPosition := substring[i] - 'a'
		if position[indexInPosition] >= 0 {
			return true
		}

		position[indexInPosition] = indexInPosition
	}

	return false
}
