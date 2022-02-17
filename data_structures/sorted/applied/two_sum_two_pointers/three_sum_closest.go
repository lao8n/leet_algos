import (
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	// choice : what to initialize currentSum to?
	// 1. initialize closestSolution to nums[i] + nums[low] + nums[high] or
	// 2. have a separate smallestDistance which initialize to maxInt
	// choice : write custom abs function for int or use float64?
	// just casted but not sure if this is kosher
	smallestDistance, closestSolution, n := math.MaxInt32, 0, len(nums)
	sort.Ints(nums)
	for i := 0; i < n-2; i++ {
		for low, high := i+1, n-1; low < high; {
			currentSum := nums[i] + nums[low] + nums[high]
			if currentSum > target {
				high--
			} else {
				low++
			}
			currentAbsDiff := int(math.Abs(float64(currentSum - target)))
			if currentAbsDiff < smallestDistance {
				closestSolution, smallestDistance = currentSum, currentAbsDiff
			}
		}
	}
	return closestSolution
}