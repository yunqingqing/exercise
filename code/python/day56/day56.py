def maxProfit(prices):
    """
    :type prices: List[int]
    :rtype: int
    """
    dp = [[0, 0]] * len(prices)
    
    for i, price in enumerate(prices):
        if i - 1 < 0:
            dp[i][0] = 0
            dp[i][1] = -price
            continue
        dp[i][1] = max(dp[i-1][1], 0-price)
        dp[i][0] = max(dp[i-1][0], dp[i-1][1]+price)
                

    print(dp)
    print(len(prices))
    return dp[len(prices)-1][0]

maxProfit([])