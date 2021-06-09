class Solution {
    public List<Integer> majorityElement(int[] nums) {
        // greedy or dynamic programming approach?
        
        // brute force
        // - O(n) time = loop over all the elements
        // - O(n) space = have a map which you count the frequency of the elements
        
        // properties of the problem
        // - only two elements can appear more than n/3 times
        
        // tricks to try
        // - if it was just alphabetic then could have fixed memory - but integers doesn't allow
        // - sorting doesn't help and would cost O(n logn) anyway
        // - sliding window / two pointers doesn't help because not try to get a combination - could have n/3 pointers
        //   but this not O(1) either
        
        // greedy approach
        // - answer is boyer-moore majority vote algorithm
        
        // 1st pass get majority leaders
        List<Integer> majorityCount = Arrays.asList(null, null, 0, 0); // [maj1, maj2, count1, count2]
        for(int i = 0; i < nums.length; i++){
            if(majorityCount.get(0) != null && nums[i] == majorityCount.get(0)) majorityCount.set(2 , majorityCount.get(2) + 1);
            else if(majorityCount.get(1) != null && nums[i] == majorityCount.get(1)) majorityCount.set(3, majorityCount.get(3) + 1);
            else if(0 == majorityCount.get(2)){ majorityCount.set(0, nums[i]); majorityCount.set(2, 1); }
            else if(0 == majorityCount.get(3)){ majorityCount.set(1, nums[i]); majorityCount.set(3, 1); }
            else{ majorityCount.set(2, majorityCount.get(2) - 1); majorityCount.set(3, majorityCount.get(3) - 1); }
        }
        // 2nd pass get actual counts
        majorityCount.set(2, 0); majorityCount.set(3, 0);
        for(int i = 0; i < nums.length; i++){
            if(nums[i] == majorityCount.get(0)) majorityCount.set(2 , majorityCount.get(2) + 1);
            else if(nums[i] == majorityCount.get(1)) majorityCount.set(3, majorityCount.get(3) + 1);
        }
        List<Integer> solutions = new ArrayList<>(2);
        if(majorityCount.get(2) > nums.length / 3) solutions.add(majorityCount.get(0));
        if(majorityCount.get(3) > nums.length / 3) solutions.add(majorityCount.get(1));
        return solutions;

        // this is quite slow - should avoid using Lists if possible
    }
}