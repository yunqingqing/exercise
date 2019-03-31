package day48

func FirstNotRepeatingChar(str string) string {
	m := make(map[rune]int, 256) // char's assic => number appear times
	runeStr := []rune(str)

	for _, char := range runeStr {
		m[char]++
	}

	for _, char := range runeStr {
		if m[char] == 1 {
			return string(char)
		}
	}

	return ""
}
