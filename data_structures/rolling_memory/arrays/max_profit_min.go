package data_structures

func maxProfit(prices []int) int {
    // could use kadane adjusted for rolling profit
    rollingMin := make([]int, len(prices))
    rollingMin[0] = prices[0]
    for i := 1; i < len(prices); i++ {
        if prices[i] < rollingMin[i - 1] {
            rollingMin[i] = prices[i]
        } else {
            rollingMin[i] = rollingMin[i - 1] 
        }
    }
    
    maxProfit := 0
    for i := 1; i < len(prices); i++ {
        if prices[i] - rollingMin[i] > maxProfit {
            maxProfit = prices[i] - rollingMin[i]
        }
    }
    return maxProfit
}