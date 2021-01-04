object Solution {
  // one pass hash map
  def twoSum(nums: Array[Int], target: Int): Array[Int] = {
    val map = scala.collection.mutable.HashMap[Int, Int]()

    for (i <- nums.indices) {
      val complement = target - nums(i)
      if (map.contains(complement))
        return Array(map(complement), i)
      else
        map.put(nums(i), i)
    }
  }

  import scala.collection.immutable.IntMap
  // recursive hash map
  def twoSum(nums: Array[Int], target: Int): Array[Int] = {
    @annotation.tailrec
    def recursiveTwoSum(map: IntMap[Int], i: Int): Array[Int] = {
      val complement = target - nums(i)
      map.find(_._1 == complement) match {
        case Some((_, complementIndex)) => Array(complementIndex, i)
        case None => recursiveTwoSum(map + (nums(i) -> i), i + 1)
      }
    }
    recursiveTwoSum(IntMap(), 0)
  }


}