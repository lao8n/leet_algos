class Solution {
    public int missingNumber(int[] nums) {
        // greedy or a dynamic programming?
        
        // brute force approach
        // test every single value from 0 to n for O(n^2)
        
        // sorting approach
        // sort values first O(n logn)
        // then go from 0 and look for missing number for O(n) processing with O(n) memory
        
        // memory approach
        // loop through adding values to set O(n)
        // then loop through values from 0 to n check set for value for O(1)
        
        // greedy approach
        // for O(n) runtime and O(1) memory we are looking at a greedy approach where we
        // calculate some sort of rolling max, min, sum, count, etc.
        // max and min dont' work unless sorted, could have multiple max mins - but hard to see how that 
        // isn't O(n)
        // rolling sum works because of gauss' formula
        int gaussSum = (nums.length + 1)* nums.length / 2;
        int sum = 0;
        for(int i = 0; i < nums.length; i++){
            sum += nums[i];
        }
        return gaussSum - sum;
    }
}