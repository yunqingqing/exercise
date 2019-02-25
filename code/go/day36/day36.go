package day36

import "fmt"


func Permutation(str string) {
	if str == "" {
		return
	}

	permutation([]rune(str), 0)
}

func permutation(str []rune, sIndex int) {
	if sIndex == len(str) {
		fmt.Printf("%s\n", string(str))
	} else {
		for i:= sIndex; i != len(str); i++ {
			str[sIndex], str[i] = str[i], str[sIndex]
			permutation(str, sIndex + 1)
			str[sIndex], str[i] = str[i], str[sIndex]
		}
	}
}
 