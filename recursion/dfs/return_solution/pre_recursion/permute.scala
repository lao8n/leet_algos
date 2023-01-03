object Solution {
    def permute(nums: Array[Int]): List[List[Int]] = {
        // brute force
        // for loop over the values 
        // n for first for loop n - 1 second n - 2 for the third etc...
        // unclear how to generate different orders
        
        // dynamic programming
        // linear recursive approach 
        // we want to build up all options by taking a different individual value for each combination
        // essentially a for loop over a slowly decreasing in size list - where each recursion is 
        // different
        
        // 1. accumulator or not? - backtracking build approach 
        // 2. linear recursion or some double approach? 
        // 3. nums is array so maybe easier with list or set
        
        def recursivePermute(numsSet : Set[Int]) : List[List[Int]] = {
            if(numsSet.isEmpty) List(List())
            else numsSet.flatMap(num => recursivePermute(numsSet - num).map(num :: _)).toList
        }
        recursivePermute(nums.toSet)
    }
}