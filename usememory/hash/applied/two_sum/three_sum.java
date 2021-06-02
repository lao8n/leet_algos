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

  // no sort approach 
  public List<List<Integer>> threeSum(int[] nums) {
    Set<List<Integer>> result = new HashSet<>();
      Set<Integer> dups = new HashSet<>();
      Map<Integer, Integer> seen = new HashMap<>();
      for (int i = 0; i < nums.length; ++i)
        if (dups.add(nums[i])) {
          for (int j = i + 1; j < nums.length; ++j) {
            int complement = -nums[i] - nums[j];
            if (seen.containsKey(complement) && seen.get(complement) == i) {
                List<Integer> triplet = Arrays.asList(nums[i], nums[j], complement);
                Collections.sort(triplet);
                result.add(triplet);
            }
            seen.put(nums[j], i);
          }
        }
      return new ArrayList(result);
  }
}