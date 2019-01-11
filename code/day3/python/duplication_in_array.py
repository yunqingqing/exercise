import unittest


def find_duplication_number_in_list(numbers : list, length : int) -> int:
    # 不符合题目要求的输入返回-1
    if length <= 0:
        return -1
    
    for i in numbers:
        if i < 0 or i > length - 1:
            return -1

    for index in numbers:
        while (index != numbers[index]):
            if numbers[index] == numbers[numbers[index]]:
                return numbers[index]
        
            # 把当前的值和下标为numbers[index]的值交换位置
            temp = numbers[index]
            numbers[index] = numbers[temp]
            numbers[temp] = temp
    return -1
    

class TestReplaceSpace(unittest.TestCase):

    def test_has_duplication_number(self):
        numbers = [2, 3, 1, 0, 5, 2, 3]
        self.assertIn(find_duplication_number_in_list(numbers, len(numbers)), [2, 3])

    def test_no_duplication_number(self):
        numbers = [2, 3, 1, 0, 5, 6, 4]
        self.assertEqual(find_duplication_number_in_list(numbers, len(numbers)), -1)

    def test_invalid_input(self):
        numbers = [1, 5, 6]
        self.assertEqual(find_duplication_number_in_list(numbers, len(numbers)), -1)

if __name__ == "__main__":
    unittest.main()