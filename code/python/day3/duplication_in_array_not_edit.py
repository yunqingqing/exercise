import unittest


def count_range(numbers : list, start : int, end : int) -> int:
    count = 0
    for num in numbers:
        if num >= start and num <= end:
            count += 1
    return count


def find_duplication_number_in_list(numbers : list, length : int) -> int:
    # 不符合题目要求的输入返回-1
    if length <= 0:
        return -1
    
    for i in numbers:
        if i < 0 or i > length - 1:
            return -1

    start = 1
    end = length -1
    while end >= start:
        middle = ((end - start) >> 1) + start
        count = count_range(numbers, start, middle)
        if start == end:
            if count > 1:
                return start
            else:
                break
            
        if count > middle - start + 1:
            end = middle
        else:
            start = middle + 1

    return -1
    

class TestReplaceSpace(unittest.TestCase):

    def test_has_duplication_number(self):
        numbers = [2, 3, 1, 6, 5, 2, 3]
        self.assertIn(find_duplication_number_in_list(numbers, len(numbers)), [2, 3])

    def test_no_duplication_number(self):
        numbers = [2, 3, 1, 7, 5, 6, 4]
        self.assertEqual(find_duplication_number_in_list(numbers, len(numbers)), -1)

    def test_invalid_input(self):
        numbers = [1, 5, 6]
        self.assertEqual(find_duplication_number_in_list(numbers, len(numbers)), -1)

if __name__ == "__main__":
    unittest.main()