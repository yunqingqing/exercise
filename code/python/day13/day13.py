"""
给你根长度为n的绳子,请把绳子剪成n段(m,n都是整数并且m,n>1)每段绳子长度记为
k[0],k[1]...k[m],请问k[0]*k[1]*...*k[m]的最大乘积是多少
例如,当n=8时,剪成2,3,3的三段,得到最大乘积18

f(n) = max(f(i) * f(n-i))
"""
import unittest


def max_product_after_cut(length):
    if length < 2:
        return 0
    elif length == 2:
        return 1
    elif length == 3:
        return 2

    products = [0] * (length+1)
    products[0] = 0 # 1
    products[1] = 1 # 2
    products[2] = 2 # 3
    products[3] = 3 # 4
    
    # 计算f(i)的最优解
    for i in range(4, length+1):
        max = 0
        for j in range(1, (i//2)+1):
            product = products[j] * products[i-j]
            if  product > max:
                max = product

            products[i] = max

    return products[-1]

class MaxProductAfterCutTestCase(unittest.TestCase):
    def test1(self):
        self.assertEqual(max_product_after_cut(1), 0)
        self.assertEqual(max_product_after_cut(2), 1)
        self.assertEqual(max_product_after_cut(3), 2)
        self.assertEqual(max_product_after_cut(4), 4)
        self.assertEqual(max_product_after_cut(5), 6)
        self.assertEqual(max_product_after_cut(6), 9)
        self.assertEqual(max_product_after_cut(7), 12)
        self.assertEqual(max_product_after_cut(50), 86093442)


if __name__ == '__main__':
    unittest.main()
