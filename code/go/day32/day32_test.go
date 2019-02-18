package day32

import "testing"

//            10
//         /      \
//        6        14
//       /\        /\
//      4  8     12  16
func TestFunc1(t *testing.T) {
	seq := []int{4, 8, 6, 12, 16, 14, 10}

	if VerifySquenceOfBST(seq, len(seq)) != true {
		t.Errorf("not correct.")
	}
}

//           5
//          / \
//         4   7
//            /
//           6
func TestFunc2(t *testing.T) {
	seq := []int{4, 6, 7, 5}

	if VerifySquenceOfBST(seq, len(seq)) != true {
		t.Errorf("not correct.")
	}
}

//               5
//              /
//             4
//            /
//           3
//          /
//         2
//        /
//       1
func TestFunc3(t *testing.T) {
	seq := []int{1, 2, 3, 4, 5}

	if VerifySquenceOfBST(seq, len(seq)) != true {
		t.Errorf("not correct.")
	}
}

// 1
//  \
//   2
//    \
//     3
//      \
//       4
//        \
//         5
func TestFunc4(t *testing.T) {
	seq := []int{5, 4, 3, 2, 1}

	if VerifySquenceOfBST(seq, len(seq)) != true {
		t.Errorf("not correct.")
	}
}

func TestFunc5(t *testing.T) {
	seq := []int{5}

	if VerifySquenceOfBST(seq, len(seq)) != true {
		t.Errorf("not correct.")
	}
}

func TestFunc6(t *testing.T) {
	seq := []int{7, 4, 6, 5}

	if VerifySquenceOfBST(seq, len(seq)) != false {
		t.Errorf("not correct.")
	}
}

func TestFunc7(t *testing.T) {
	seq := []int{4, 6, 12, 8, 16, 14, 10}

	if VerifySquenceOfBST(seq, len(seq)) != false {
		t.Errorf("not correct.")
	}
}
