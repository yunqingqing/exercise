def maxProfit(prices):
    """
    :type prices: List[int]
    :rtype: int
    """
    dp = [[0, 0] for i in prices] 
    for i, price in enumerate(prices):
        if i - 1 < 0:
            dp[i][0] = 0
            dp[i][1] = -price
            continue
        dp[i][1] = max(dp[i-1][1], 0-price)
        dp[i][0] = max(dp[i-1][0], dp[i-1][1]+price)
                
        print(i)
        print(dp)
    return dp[len(prices)-1][0]

print(maxProfit([7,1,5,3,6,4]))