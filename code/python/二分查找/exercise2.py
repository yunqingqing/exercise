""" 寻找比目标字母大的最小字母
给定一个只包含小写字母的有序数组letters 和一个目标字母 target，寻找有序数组里面比目标字母大的最小字母。

数组里字母的顺序是循环的。举个例子，如果目标字母target = 'z' 并且有序数组为 letters = ['a', 'b']，则答案返回 'a'。

示例:

输入:
letters = ["c", "f", "j"]
target = "a"
输出: "c"

输入:
letters = ["c", "f", "j"]
target = "c"
输出: "f"

输入:
letters = ["c", "f", "j"]
target = "d"
输出: "f"

输入:
letters = ["c", "f", "j"]
target = "g"
输出: "j"

输入:
letters = ["c", "f", "j"]
target = "j"
输出: "c"

输入:
letters = ["c", "f", "j"]
target = "k"
输出: "c"
""" 
"c", "c", "c", "c", "c", "f", "j"   "a"
# 0    1     2    3    4    5    6
#            x      x
import unittest

def next_greatest_letter(letters, target):
    l, r = 0, len(letters) - 1
    result = letters[0]
    while l <= r:
        mid = (l+r)//2
        tmp = letters[mid]
        if tmp > target:
            r = mid - 1
            result = tmp
        else:
            l = mid + 1
    return result

class NextGreatestLetterTestCase(unittest.TestCase):
    def test_next_greatest_letter(self):
        self.assertEqual(next_greatest_letter(["c", "f", "j"], "a"), "c")
        self.assertEqual(next_greatest_letter(["c", "f", "j"], "d"), "f")
        self.assertEqual(next_greatest_letter(["c", "c", "c", "c", "f", "j"], "c"), "f")
        self.assertEqual(next_greatest_letter(["c", "f", "j", "j", "j", "j"], "g"), "j")
        self.assertEqual(next_greatest_letter(["c", "f", "j"], "j"), "c")
        self.assertEqual(next_greatest_letter(["c", "f", "j"], "k"), "c")

    
if __name__ == "__main__":
    unittest.main()
