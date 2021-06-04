class Solution {
    // greedy or a dynamic programming?
    // use dynamic = four-sum = 2-sum + two additional for loops
    
    // two-sum can be solved in 3 ways
    // 1. brute force 2 for loops
    // 2. sort then two pointers
    // 3. hash map
    
    // because we want to avoid duplicates more easily and the cost of sorting it small
    // we will go with two-pointers approach
    // i have for my scala implementation pulled off a recursive binary search using hash
    // maps but here we will practice using more standard for loops
    public List<List<Integer>> fourSum(int[] nums, int target) {
        Arrays.sort(nums);
        Set<List<Integer>> solutions = new HashSet<>();
        for(int i = 0; i < nums.length - 3; i++){
            for(int j = i + 1; j < nums.length - 2; j++){
                twoSum(nums, target, i, j, j + 1, nums.length - 1, solutions);
            }
        }
        return new ArrayList<>(solutions);
    }
    
    private void twoSum(int[] nums, int target, int firstIndex, int secondIndex, int leftIndex, int rightIndex, Set<List<Integer>> solutions){
        // assume nums is sorted
        while(leftIndex < rightIndex){
            int sum = nums[firstIndex] + nums[secondIndex] + nums[leftIndex] + nums[rightIndex];
            if(sum < target) leftIndex++;
            else if(sum > target) rightIndex--;
            else if(sum == target){
                solutions.add(Arrays.asList(nums[firstIndex], nums[secondIndex], nums[leftIndex], nums[rightIndex]));
                leftIndex++;
                rightIndex--;
            }
        }
    }
    
    
}