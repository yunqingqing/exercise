"""
Input:5
Output: True
Explanation: 1*1 + 2*2
"""

def judgeSquareSum(number):
    i, j = 0, number
    while i <= j:
        now = i*i + j*j
        if now == number:
            return True
        elif now < number:
            j -= 1
        else:
            i += 1

    return False