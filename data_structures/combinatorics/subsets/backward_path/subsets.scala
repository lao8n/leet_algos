object Solution {
    def subsets(nums: Array[Int]): List[List[Int]] = {
        // brute force
        // greedy approach
        // dynamic programming
        // - memoization - backtracking etc. 
        // key problem is how do we append to the list? 
        // problem how should i recurse so that i 1. avoid duplicates and 2. get all the combinations?
        nums match {
            case Array() => List(List[Int]())
            case nums => 
                val tailSubsets = subsets(nums.tail)
                tailSubsets ::: tailSubsets.map(sub => nums.head :: sub)
        }
    }
}