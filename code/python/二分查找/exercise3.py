"""有序数组中的单一元素
给定一个只包含整数的有序数组，每个元素都会出现两次，唯有一个数只会出现一次，找出这个数。

示例 1:

输入: [1,1,2,3,3,4,4,8,8]
输出: 2
示例 2:

输入: [3,3,7,7,10,11,11]
输出: 10
注意: 您的方案应该在 O(log n)时间复杂度和 O(1)空间复杂度中运行。
"""

import unittest

def single_non_duplicate(nums): 
    l, r = 0, len(nums)-1

    while l <= r:
        mid = (l +r) // 2
        mid_num = nums[mid]
        if mid_num not in nums[l:mid] and mid_num not in nums[mid:r]:
            return mid_num
        elif mid_num not in nums[l:mid]:
            r = mid
        else:
            l = mid + 1

    return 

class SingleNonDuplicateTestCase(unittest.TestCase):
    def test_single_non_duplicate(self):
        self.assertEqual(single_non_duplicate([1,1,2,3,3,4,4,8,8]), 2)
        self.assertEqual(single_non_duplicate([3,3,7,7,10,11,11]), 10)


if __name__ == '__main__':
    unittest.main()