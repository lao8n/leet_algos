package data_structures

import "math"

// interdependent problems combined for a solution -> dynamic programming
// however can use non-standard approach by merging sequences - exploiting price info which allows greedier approach
// time = O(n(n-k)) if 2k <= n and O(n) if 2k > n where n is the length of the price seq
// space = O(n) for list of transactions
func maxProfit(k int, prices []int) int {
	if len(prices) == 0 || k == 0 {
		return 0
	}

	// find all consecutively increasing subsequences
	transactions := make([]incrSub, 0)
	bp, sp := math.MaxInt, -1
	// at peak [4, 5, 3] -> sp = 4 -> 5
	// at trough [3, 1, 2] -> bp
	for _, p := range append(prices, 0) {
		// reset
		if p < sp {
			transactions = append(transactions, incrSub{buyPrice: bp, sellPrice: sp})
			bp, sp = math.MaxInt, -1
		}
		// if p decreasing buy, increasing sell
		if p < bp {
			bp = p
		} else if p > sp {
			sp = p
		}
	}

	for len(transactions) > k {
		// choice either delete or merge
		minDeleteLoss, minMergeLoss := math.MaxInt, math.MaxInt
		minDeleteLossIndex, minMergeLossIndex := 0, 0
		for i, t := range transactions {
			// delete loss
			tProfit := t.sellPrice - t.buyPrice
			if tProfit < minDeleteLoss {
				minDeleteLoss = tProfit
				minDeleteLossIndex = i
			}
			// merge loss
			// before merge (t1.s - t1.b) + (t2.s - t2.b)
			// after merge (t2.s - t1.s)
			// diff (t1.s - t2.b)
			if i > 0 {
				tMergedProfit := transactions[i-1].sellPrice - transactions[i].buyPrice
				if tMergedProfit < minMergeLoss {
					minMergeLoss = tMergedProfit
					minMergeLossIndex = i // index of second
				}
			}
		}
		// choice either delete or merge - can we do so greedily?
		if minMergeLoss < minDeleteLoss {
			mergedSub := incrSub{
				buyPrice:  transactions[minMergeLossIndex-1].buyPrice,
				sellPrice: transactions[minMergeLossIndex].sellPrice,
			}
			transactions = append(append(transactions[:minMergeLossIndex-1], mergedSub), transactions[minMergeLossIndex+1:]...)
		} else {
			transactions = append(transactions[:minDeleteLossIndex], transactions[minDeleteLossIndex+1:]...)
		}
	}

	cumulativeProfit := 0
	for _, t := range transactions {
		cumulativeProfit += t.sellPrice - t.buyPrice
	}
	return cumulativeProfit
}

type incrSub struct {
	buyPrice  int
	sellPrice int
}
