"""
Input: number={2，7，11，15}， target=9
Output: index1=1, index2=2
"""

def two_sum(numbers, target):
    i, j = 0, len(numbers)

    while i != j:
        now = numbers[i] + numbers[j]
        if now == target:
            return i, j
        elif now < target:
            i += 1
        else:
            j += 1

    return 