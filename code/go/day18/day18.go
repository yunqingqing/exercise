package day18

import "fmt"

func Math(str, pattern string) bool {
	pStr := 0
	pPattern := 0
	return matchCore(str, pattern, pStr, pPattern, len(str), len(pattern))
}

func matchCore(str, pattern string, pStr, pPattern, strLen, patternLen int) bool {
	if strLen == pStr && pPattern == patternLen{
		return true
	}

	if pStr != strLen && pPattern > patternLen {
		return false
	}

	// 防止array越界，单独判断取值
	var patternChar, strChar string
	if pPattern >= patternLen {
		patternChar = ""	
	} else {
		patternChar = string(pattern[pPattern])
	}
	if pStr >= strLen {
		strChar = ""	
	} else {
		strChar = string(str[pStr])
	}

	if (pPattern + 1) < patternLen && string(pattern[pPattern+1]) == "*" {

		if patternChar == strChar || (patternChar == "." && strLen != pStr) {
			// 进入下一个状态
			enterNext := matchCore(str, pattern, pStr+1, pPattern+2, strLen, patternLen)
			// 留在当前状态，继续匹配*
			stay := matchCore(str, pattern, pStr+1, pPattern, strLen, patternLen)
			// 略过一个*号
			ignore := matchCore(str, pattern, pStr, pPattern+2, strLen, patternLen)

			return enterNext || stay || ignore
		} else {
			// 略过一个*号
			return matchCore(str, pattern, pStr, pPattern+2, strLen, patternLen)
		}
	}

	if patternChar == strChar || (patternChar == "." && strLen != pStr) {
		return matchCore(str, pattern, pStr+1, pPattern+1, strLen, patternLen)
	}

	return false
}
