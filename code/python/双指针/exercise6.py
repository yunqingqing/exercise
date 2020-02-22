"""
Input:
s = "abpcplea", d = ["ale","apple","monkey","plea"]

Output:
"apple"
给定一个字符串和一个字符串字典，找到字典里面最长的字符串，
该字符串可以通过删除给定字符串的某些字符来得到。如果答案不止一个，
返回长度最长且字典顺序最小的字符串。如果答案不存在，则返回空字符串。
"""

import unittest

def find_longest_word(string, d):
    longest_word = ""

    for word in d:
        if is_sub_str(string, word) and len(word) > len(longest_word):
            longest_word = word

    return longest_word


def is_sub_str(s, target):
    i, j = 0, 0
    while i<len(s) and j < len(target):
        if s[i] == target[j]:
            i += 1
            j += 1
        else:
            i += 1
    return j == len(target)


class FindLongestWordTestCase(unittest.TestCase):
    def test_find_longest_word(self):
        self.assertEqual(find_longest_word("abpcplea", ["ale","apple","monkey","plea"]), "apple")

    def test_sub_str(self):
        self.assertEqual(is_sub_str("abpcplea", "apple"), True)

if __name__ == "__main__":
    unittest.main()
