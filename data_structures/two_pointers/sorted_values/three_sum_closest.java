class Solution {
    public int threeSumClosest(int[] nums, int target) {
        // greedy or dynamic programming? dynamic because can treat as similar to two-sum problem
        // single answer = do not need to accumulate solutions and can return immediately
        // closest to target = either 1. sort the array and choose the closest 2. retain in memory the closest
        
        // two sum problem
        // 1. brute force two for loops
        // 2. hash map
        // 3. sort first then two pointers
        Arrays.sort(nums);
        int closestDifference = Integer.MAX_VALUE;
        for(int i = 0; i < nums.length; i++){
            closestDifference = twoSum(nums, target, i, closestDifference);
        }
        return target - closestDifference;
    }
    
    // key thing is cannot use closestSum because then start with Integer.MAX_VALUE and can get stuck
    // need to use closestDifference
    private int twoSum(int[] nums, int target, int numIndex, int closestDifference){
        // assume nums is sorted
        int leftIndex = numIndex + 1;
        int rightIndex = nums.length - 1;
        while(leftIndex < rightIndex){
            int sum = nums[numIndex] + nums[leftIndex] + nums[rightIndex];
            if(Math.abs(target - sum) < Math.abs(closestDifference)){
                closestDifference = target - sum;
            }
            if(sum == target){
                return 0;
            }
            else if(target - sum > 0){
                leftIndex++;
            }
            else { // target - sum < 0
                rightIndex--;
            }
        }
        return closestDifference;
    }
}