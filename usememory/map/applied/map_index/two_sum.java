// brute force
public int[] twoSum(int[] nums, int target){
  for (int i = 0; i < nums.length; i++){
    for (int j = i + 1; j < nums.length; j++){
      if (nums[j] == target - nums[i]){
        return new int[] { i , j };
      }
    }
  }
  throw new IllegalArgumentException("no two sum solution");
}

// two pass hash map
public int[] twoSum(int[] nums, int target){
  Map<Integer, Integer> numsMap = new HashMap<>();
  for (int i = 0; i < nums.length; i++){
    numsMap.put(nums[i], i);
  } 
  for (int i = 0; i < nums.length; i++){
    int complement = target - nums[i];
    // need to check not getting the same value twice i.e. if target
    // is 8 we get 4 and 8's complement 4 again
    if (numsMap.containsKey(complement) && numsMap.get(complement) != i){
      return new int[] { numsMap.get(complement), i};
    }
  }
  throw new IllegalArgumentException("no two sum solution");
}

// one pass hash map
public int[] twoSum(int[] nums, int target){
  Map<Integer, Integer> numsMap = new HashMap<>();
  for (int i = 0; i < nums.length; i++){
    int complement = target - nums[i];
    if (numsMap.containsKey(complement) && numsMap.get(complement) != i){
      return new int[] { numsMap.get(complement), i};
    }
    numsMap.put(nums[i], i);
  }
  throw new IllegalArgumentException("no two sum solution");
}