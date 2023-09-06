object Solution {
    def threeSumClosest(nums: Array[Int], target: Int): Int = {
        // q: greedy or a dynamic programming? dynamic = two sum problem ++
        
        // two sum 
        // 1. brute force w/ for loops
        // 2. map to track location of values 
        // 3. sort the array + two-pointers -> best choice
        
        // only need one solution so don't need to have solutions accumulator etc.
        // scala choice: immutable data structures & recursive function calls or
        // mutable data structures & foldLeft? 
        // if foldLeft need to keep track of two indices and have if statement on leftIndex < rightIndex
        // cleaner to go for recursive functional calls
        
        // not pure dynamic programming = because we need to test each sum (including num) against target
        val sortedNums = nums.sorted
        
        def recursiveThreeSum(numIndex: Int, leftIndex: Int, rightIndex: Int, target: Int, closestDiff: Int): Int = {
            val sum = sortedNums(numIndex) + sortedNums(leftIndex) + sortedNums(rightIndex)
            val diff = target - sum
            val newClosestDiff = if(diff.abs < closestDiff.abs) diff else closestDiff
            if(sum < target && leftIndex + 1 < rightIndex)
                recursiveThreeSum(numIndex, leftIndex + 1, rightIndex, target, newClosestDiff)
            else if(sum > target && leftIndex < rightIndex - 1)
                recursiveThreeSum(numIndex, leftIndex, rightIndex - 1, target, newClosestDiff)
            else if(numIndex + 1 < sortedNums.size - 2) // need room for leftIndex and rightIndex
                recursiveThreeSum(numIndex + 1, numIndex + 2, sortedNums.size - 1, target, newClosestDiff)
            else
                target - newClosestDiff
        }
        recursiveThreeSum(0, 1, sortedNums.size - 1, target, Int.MaxValue)
    }
}