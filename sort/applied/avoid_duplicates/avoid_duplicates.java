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

  // this is the same as the two_sum one pass hash map
  // difference is that we use a set because the condition we want to test
  // with map lookup index can never happen
  // int complement = target - nums[i];
  // if (numsMap.containsKey(complement) && numsMap.get(complement) != i){
  void twoSum(int[] nums, int targetIndex, List<List<Integer>> result){
    Set<Integer> numsSet = new HashSet<>();
    for (int i = targetIndex + 1; i < nums.length; i++){
      int complement = -nums[targetIndex] - nums[i];
      if(numsSet.contains(complement)){
        result.add(Arrays.asList(nums[targetIndex], nums[i], complement));
        // loop through duplicates
        while(i + 1 < nums.length && nums[i] == nums[i + 1])
          i++;
      }
      numsSet.add(nums[i]);
    }
  }
}