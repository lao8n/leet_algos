package data_structures

import "math"

// clarifying questions
// * 1-indexed array where prices[i] is the number of coins to purchase the ith fruit?
// * purchase ith fruit get the next i fruits for free
// * minimum number of coins - to purchase lowest price?

// thoughts
// * can we re-write the costs as array after you have bought? e.g. prices in example 2 becomes [0, 0, 1, 1]
// or in example 1 [0, 0, 2] - although still optimal to buy before hand..
// * basically when you purchase you are buying a set of fruits - and you want their combined overlap to be all the fruits.. i.e. their intersection should equal the entire array..
// * always need to purchase first fruit - but then we have the same scenario again but with two choices?

// approaches
// * not just pick the minimum fruits because also need to think about how many fruits you get for that..
// * surely this is a dynamic programming problem? typically with minimum coin problems you update the remaining total but here we have an entire array to update
// * for each number can you just greedily pick the first number? probably not right? because even if locally optimal but get a lot of other with it... must be dp
// * is it a problem we only consider even numbers going forward and dp array has half zeroes?
func minimumCoins(prices []int) int {
	// state is value of coins and all coins that are in intersection
	// dp := make([]int, len(prices)) // dp[i] is where we have gotten up to in the coins array - can either go forward or go backwards
	dp := make([]int, len(prices)+1) // make handling 1st index easier
	for i := range dp {
		if i != 0 {
			dp[i] = math.MaxInt
		}
	}

	for i, p := range prices {
		// go forwards - all ways to use p - dp represents single pass forward best way to get to jth index
		for j := i + 1; j < min(len(prices), 2*(i+1))+1; j++ {
			dp[j] = min(dp[j], dp[i]+p)
		}
	}
	return dp[len(prices)]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// func recursiveMinCoins(prices []int, dp []int, i int, rollingCost int) int {
//     if i > len(prices) {
//         return rollingCost
//     }
//     if dp[i - 1] != 0 {
//         return dp[i - 1]
//     }
//     // consider all possible ways to get the ith value
//     // buy ith - and then advance i + i..
//     // buy i-1 and advance i - 1...
//     lowestCost := math.MaxInt
//     j := 0
//     // go backwards - all the ways we can pay for the ith fruit..
//     for 2 * (i - j) >= i { // if purchase the ith - jth fruit we can get the next ith - j fruits for free
//         currentCost := recursiveMinCoins(prices, dp, 2 * (i - j) + 1, rollingCost + prices[i - j - 1])
//         if currentCost < lowestCost { // maybe this is too greedy?
//             lowestCost = currentCost
//         }
//         j++
//     }
//     dp[i - 1] = lowestCost
//     return lowestCost
// }

// [38,23,27,32,47,45,48,24,39,26,37,42,24,45,27,26,15,16,26,6]
// Output:
// 134
// Expected:
// 132
// Stdout:
// [134 0 134 0 134 0 134 0 134 0 134 0 134 0 137 0 152 0 167 0]
