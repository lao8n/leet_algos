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

  import scala.collection.immutable.IntMap
  // recursive hash map
  def twoSum(nums: Array[Int], target: Int): Array[Int] = {
    @annotation.tailrec
    def recursiveTwoSum(numsMap: IntMap[Int], i: Int): Array[Int] = {
      val complement = target - nums(i)
      numsMap.find(_._1 == complement) match {
        case Some((_, complementIndex)) => Array(complementIndex, i)
        case None => recursiveTwoSum(numsMap + (nums(i) -> i), i + 1)
      }
    }
    recursiveTwoSum(IntMap(), 0)
  }
}