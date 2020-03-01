"""
实现 int sqrt(int x) 函数。

计算并返回 x 的平方根，其中 x 是非负整数。

由于返回类型是整数，结果只保留整数的部分，小数部分将被舍去。

示例 1:

输入: 4
输出: 2
示例 2:

输入: 8
输出: 2
说明: 8 的平方根是 2.82842..., 
     由于返回类型是整数，小数部分将被舍去。
"""
import unittest

def my_sqrt(num):
    l, r = 1, num
    while l <= r:
        mid = (l + r) // 2
        if mid * mid < num:
            l = mid + 1
        elif mid * mid > num:
            r = mid - 1
        else:
            return mid

    return r


class MySqrtTestCase(unittest.TestCase):
    def test_my_sqrt(self):
        self.assertEqual(my_sqrt(4), 2)

    def test_my_sqrt1(self):
        self.assertEqual(my_sqrt(8), 2)


if __name__ == '__main__':
    unittest.main()