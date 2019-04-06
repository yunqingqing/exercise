package day49

import "fmt"

func InversePairs(data []int, length int) int {
	if length < 0 {
		return 0
	}

	copyData := make([]int, length)
	
	for i:= 0; i < length; i++ {
		copyData[i] = data[i]
	}

	count := InversePairesCore(data, copyData, 0, length - 1)
	return count
}

func InversePairesCore(data, copyData []int, start, end int) int {
	if start == end {
		copyData[start] = data[start]
		return 0
	}

	length := (end - start) / 2
	left := InversePairesCore(copyData, data, start, start + length)
	right := InversePairesCore(copyData, data, start + length + 1, end)

	i := start + length
	j := end
	indexCopy := end
	count := 0
	fmt.Println("dsfadfa~~~~length", length)
	fmt.Println("dsfadfa~~~~start", start)
	fmt.Println("dsfadfa~~~~end", end)
	fmt.Println("dsfadfa=>copyData", copyData)
	fmt.Println("dsfadfa=>data", data)
	fmt.Println("dsfadfa~~~~i", i)
	for i >= start && j >= start + length + 1 {
		indexCopy--
		if data[i] > data[j] {
			i--
			fmt.Println("dsfadfa=>indexCopy", indexCopy)
			fmt.Println("dsfadfa=>i", i)
			copyData[indexCopy] = data[i]
			count += j - start - length
		} else {
			j--
			fmt.Println("dsfadfa=>indexCopy", indexCopy)
			fmt.Println("dsfadfa=>j", j)
			copyData[indexCopy] = data[j]
		}
	}
	
	fmt.Println("dsfadfa=>copyData", copyData)
	fmt.Println("dsfadfa=>data", data)
	for ; i >= start; i-- {
		indexCopy--
		fmt.Println("dsfadfa=>indexCopy", indexCopy)
		fmt.Println("dsfadfa=>i", i)
		copyData[indexCopy] = data[i]
	}
	for ; j >= start + length +  1; j-- {
		indexCopy--
		fmt.Println("dsfadfa=>indexCopy", indexCopy)
		fmt.Println("dsfadfa=>j", j)
		copyData[indexCopy] = data[j]
	}

	return left + right +count
}