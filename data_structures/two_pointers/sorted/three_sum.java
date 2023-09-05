class Solution{
  public List<List<Integer>> threeSum(int[] nums){
    // sort array to make looping through duplicates easier
    Arrays.sort(nums);
    List<List<Integer>> result = new ArrayList<>();
    for (int targetIndex = 0; targetIndex < nums.length && nums[targetIndex] <= 0; targetIndex++){
      if(targetIndex == 0 || nums[targetIndex - 1] != nums[targetIndex]){
        twoSum(nums, targetIndex, result);
      }
    }
    return result;
  }

  void twoSum(int[] nums, int targetIndex, List<List<Integer>> result){
    int lo = targetIndex + 1, hi = nums.length - 1;
      while (lo < hi) {
        int sum = nums[targetIndex] + nums[lo] + nums[hi];
        if (sum < 0) {
          ++lo;
        } else if (sum > 0) {
          --hi;
        } else {
          result.add(Arrays.asList(nums[targetIndex], nums[lo++], nums[hi--]));
          while (lo < hi && nums[lo] == nums[lo - 1])
            ++lo;
        }
  }
}