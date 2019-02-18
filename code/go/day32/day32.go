package day32

func VerifySquenceOfBST(sequence []int, length int) bool {
	if sequence == nil && length <= 0 {
		return false
	}

	root := sequence[length-1]

	var i, j int
	for i = 0; i < length-1; i++ {
		if sequence[i] > root {
			break
		}
	}

	// 在二叉搜索树中右子树的结点大于根结点
	for j = i; j < length-1; j++ {
		if sequence[j] < root {
			return false
		}
	}

	left := true
	if i > 0 {
		// 判断左子树是否为bst
		left = VerifySquenceOfBST(sequence, i)
	}

	right := true
	if i < length-1 {
		// 判断右子树是否为bst
		right = VerifySquenceOfBST(sequence[i:length-1], length-i-1)
	}

	return left && right
}
