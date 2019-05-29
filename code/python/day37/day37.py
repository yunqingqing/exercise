"""
3, 5, 2, 1, 8, 6
选取1作为基准: 3, 5, 2, 6, 8, 1
"""
import unittest
import random


def swap(l, index1, index2):
    l[index1], l[index2] = l[index2], l[index1]


def partition(numbers, start, end):
    parviot = random.randint(start, end-1)
    small = start - 1
    swap(numbers, parviot, end)

    for i in range(end-start+1):
        index = start+i
        if numbers[index] < numbers[end]:
            small += 1
            if index != small:
                swap(numbers, index, small)

    small += 1
    swap(numbers, small, end)

    return small


def quick_sort(numbers, start, end):
    if start == end:
        return

    privot = partition(numbers, start, end)
    if privot > start:
        quick_sort(numbers, start, privot-1)
    if end > privot:
        quick_sort(numbers, privot+1, end)

    return numbers



class TestQuickSort(unittest.TestCase):

    def test_sort(self):
        self.assertEqual(quick_sort([3, 5, 2, 1, 8, 6], 0, 5), [1, 2, 3, 5, 6, 8])

if __name__ == '__main__':
    unittest.main()
