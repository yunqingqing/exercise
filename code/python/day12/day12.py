
"""请实现一个函数,输入一个整数,输出该数二进制表示中的1的个数.
"""
import unittest


def number_of_1(num):
    count = 0

    while num:
        count += 1
        num = (num-1) & num

    return count


class NumberOf1TestCase(unittest.TestCase):
    def test_1(self):
        self.assertEqual(number_of_1(0), 0)
        self.assertEqual(number_of_1(1), 1)
        self.assertEqual(number_of_1(10), 2)
        self.assertEqual(number_of_1(0x7FFFFFFF), 31)
        self.assertEqual(number_of_1(0xFFFFFFFF), 32)


if __name__ == '__main__':
    unittest.main()
