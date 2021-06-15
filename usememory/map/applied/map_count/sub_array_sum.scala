object Solution {
    def subarraySum(nums: Array[Int], k: Int): Int = {
        // brute force
        // for loop over each starting index
        // for loop over each ending index
        // for loop to calculate the sum
        // increment for each sum that matches
        // wasted computation because adding the same numbers again and again -> dynamic programming memoization
        // so O(n^3)
        
        // sorting
        // sort O(n logn) 
        
        // memory
        // constant memory thing where we keep rolling sum - two pointers start and end index
        // O(n), memory O(1)
        // turns out cannot do sliding window because can have negative numbers
        // does it work if we sorted?
        
        // greedy
        // min, max, sum -> doesn't help
        // do rolling cumulative sum - where if difference between two sums is k
        // we have a hashmap with sum -> number of occurences of sum
        // check if map contains (sum - k)
        nums.foldLeft(Map(0 -> 1), 0, 0){
            case ((sumCountMap, sum, count), num) =>
                val newSum = sum + num
                (sumCountMap + (newSum -> (sumCountMap.getOrElse(newSum, 0) + 1)),
                 newSum, 
                 count + sumCountMap.getOrElse(newSum - k, 0))
        }._3
        
        
        // dynamic programming
        // for loop for rolling cumulative sum
        // then calculate difference in sums with start and end index
        // however because we have negative numbers we cannot do sliding window
        // so still O(n^2)
        // or just do a rolling sum for each end index - still O(n^2)
    }
}