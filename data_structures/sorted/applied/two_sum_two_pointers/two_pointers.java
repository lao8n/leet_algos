class Solution {
  public int[] twoSum(int[] nums, int target) {
    int startIndex = 0, endIndex = nums.length - 1;
    while(startIndex < endIndex){
      int sum = nums[startIndex] + nums[endIndex];
      if(sum == target){
          return new int[]{startIndex + 1, endIndex + 1};
      }
      else if (sum < target){
          startIndex++;
      }
      else{
          endIndex--;
      }
    }
    return new int[]{};
  }
}