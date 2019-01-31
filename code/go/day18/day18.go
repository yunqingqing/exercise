package day18

func Math(str, pattern string) bool {
	pStr := 0
	pPattern := 0
	return matchCore(str, pattern, pStr, pPattern, len(str), len(pattern))
}

func matchCore(str, pattern string, pStr, pPattern, strLen, patternLen int) bool {
	if strLen == pStr && pPattern >= patternLen{
		return true
	}

	if strLen != strLen && pPattern >= patternLen {
		return false
	}

	if string(pattern[pPattern+1]) == "*" {
		if string(pattern[pPattern]) == string(str[pStr]) {
			return matchCore(str, pattern, pStr+1, pPattern+2, strLen, patternLen) || matchCore(str, pattern, pStr+1, pPattern, strLen, patternLen) || matchCore(str, pattern, pStr, pPattern+2, strLen, patternLen)
		} else {
			return matchCore(str, pattern, pStr, pPattern+2, strLen, patternLen)
		}
	}

	if string(pattern[pPattern]) == string(str[pStr]) && string(pattern[pPattern]) == "." && strLen != pStr {
		return matchCore(str, pattern, pStr+1, pPattern+1, strLen, patternLen)
	}

	return false
}
