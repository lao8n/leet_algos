object Solution {
    def majorityElement(nums: Array[Int]): List[Int] = {
        // brute force
        // for loop over all the elements with a separate Map for elements and count
        // then for loop over values in map checking if > n/3 and outputting as an array
        
        // observations
        // a maximum of 2 elements can appear over n/3 times 
        // can't sort because that doesn't give us linear time
        // can probably stop early when we know that n/3 times cannot exist
        // so you can't avoid going through the whole array in case there is a element right on the threshold
        
        // greedy trick
        // there can only be max 2 with > n/3 occurrences
        // apparently you then only need to keep track of 2 top candidates and 2 alternatives - not sure why specifically this is sufficient
        // then need a second pass to actual calculate if it crosses the threshold
        
        // 1st pass determine majority leaders
        nums.foldLeft(Array.ofDim[Int](4)){ // [maj1, maj2, count1, count2]
            case (majCount, num) =>
                if(majCount(0) == num) majCount.updated(2, majCount(2) + 1)
                else if(majCount(1) == num) majCount.updated(3, majCount(3) + 1)
                else if(majCount(2) == 0) majCount.updated(0, num).updated(2, 1)
                else if(majCount(3) == 0) majCount.updated(1, num).updated(3, 1)
                else majCount.updated(2, majCount(2) - 1).updated(3, majCount(3) - 1)
        }
        // rather than second pass cleaner to use filter
        .take(2)
        .distinct
        .filter(n => nums.count(_ == n) > nums.length / 3)
        .toList
    }
}