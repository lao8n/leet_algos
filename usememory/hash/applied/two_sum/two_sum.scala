object Solution {
  // one pass hash map
  def twoSum(nums: Array[Int], target: Int): Array[Int] = {
    val numsMap = scala.collection.mutable.HashMap[Int, Int]()

    for (i <- nums.indices) {
      val complement = target - nums(i)
      if (numsMap.contains(complement))
        return Array(numsMap(complement), i)
      else
        numsMap.put(nums(i), i)
    }
  }

  // one line 
  def twoSum(nums: Array[Int], target: Int): Array[Int] = {
    nums.zipWithIndex
      .foldLeft(Map[Int, Int]()){
        case (numsMap, numsTuple @(x, complementIndex)) => 
          numsMap.get(target - x)
                 .map(index => return Array(index, complementIndex))
                 .getOrElse(numsMap + numsTuple)
      }
  }

  import scala.collection.immutable.HashMap
  def twoSum(nums: Array[Int], target: Int): Array[Int] = {
    @annotation.tailrec
    def recursiveTwoSum(map: HashMap[Int, Int], i: Int): Array[Int] = {
      val complement = target - nums(i)
      map.get(complement) match {
        case Some(complementIndex) => Array(complementIndex, i)
        case None => recursiveTwoSum(map + (nums(i) -> i), i + 1)
      }
    }
    recursiveTwoSum(HashMap(), 0)
  }
}