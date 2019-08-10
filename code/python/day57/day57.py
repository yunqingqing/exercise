import unittest

def partition(nums, l, r):
    i=l+1
    j=r
    
    while 1:
        while i < r and nums[i] <= nums[l]:
            i += 1
        while j > l and nums[j] >= nums[l]:
            j -= 1
        if i>=j:
            break  
        nums[i], nums[j] = nums[j], nums[i]
        
    nums[l], nums[j] = nums[j], nums[l]
    return j


def findKthLargest(nums, k):
    """
    :type nums: List[int]
    :type k: int
    :rtype: int
    """
    l=0
    k = len(nums) - k
    h=len(nums) -1
    
    while l <= h:
        j = partition(nums, l, h)
        if j == k:
            break
        elif j < k:
            l = j + 1
        else:
            h = j -1
        
    return nums[j]


class FindKthLargestTestCase(unittest.TestCase):
    def test_find_kth_largest(self):
        self.assertEqual(findKthLargest([3,2,3,1,2,4,5,5,6], 4), 4)
        self.assertEqual(findKthLargest([3,2,1,5,6,4], 2), 5)
        self.assertEqual(findKthLargest([1], 1), 1)
        self.assertEqual(findKthLargest([3,3,3,3,3,3,3,3,3], 1), 3)


if __name__ == '__main__':
    unittest.main() 