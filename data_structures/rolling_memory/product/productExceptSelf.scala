object Solution {
    def productExceptSelf(nums: Array[Int]): Array[Int] = {
        // brute force
        // two for loops one for each answer and another to loop through and multiply the product
        
        // with division
        // single loop through to get the product of all 
        // then second loop dividing by nums[i] or multiplying by 1/nums[i]
        
        // with ring
        // have an array of n length and multiply each element by the next element (wrapping around) n-1 times but this is n * n-1
        
        // greedy or a dynamic programming approach?
        // greedy approach where you maintain two products - one with nums[i] 
        // and one without, but somehow you can drop one 
        // dynamic approach where you can share computation for each i and i  + 1
        // most of the computation is the same except two elements so maybe you can do a rolling recursive thing 
        
        // if you consider the answers we want to have for
        // [x1, x2, x3, x4]
        // answer is
        // [x2x3x4, x1x3x4, x1x2x4, x1x2x3]
        // so two rolling arrays
        // [0, x1, x1x2, x1x2x3]
        // [0, x4, x3x4, x2x3x4]
        
        // complexity O(n) [1, x0, x0x1, ...]
        
        // 
        // val leftOfI = nums.scanLeft(1)(_*_).dropRight(1)
        // val rightOfI = nums.scanRight(1)(_*_).tail
        // leftOfI zip rightOfI map {case (l,r) => l*r}
        // constant time 
        nums.scanLeft(1)(_*_).dropRight(1) // [1, x0, x0x1, ..., x0x1...xn-1]
            .zipWithIndex // [(1, 0), (x0, 1), (x0x1, 2),...(x0x1...xn-1, n-1)]
            .scanRight((1, 1)){case ((l, i), (a, r)) => (l * r, nums(i) * r)} // [(x1...xn-1, x0...xn-1), (x0x2...xn-1, x1...xn-1), (x0x1x3...xn-1, x2...xn-1).., (x0x1...xn-1, 1),(1,1)]
            .map{case (a, _) => a} // [x1...xn-1, x0x2...xn-1, x0x1x3...xn-1.., x0x1...xn-1, 1]
            .dropRight(1) // [x1...xn-1, x0x2...xn-1, x0x1x3...xn-1.., x0x1...xn-1]
    }
}