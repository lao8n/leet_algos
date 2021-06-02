object Solution {
  def twoSum(numbers: Array[Int], target: Int): Array[Int] = {
    @annotation.tailrec
    def recursiveTwoSum(startIndex: Int, endIndex: Int): Array[Int] = {
        val sum = numbers(startIndex) + numbers(endIndex)
        if(sum < target && endIndex - startIndex > 1) recursiveTwoSum(startIndex + 1, endIndex)
        else if(sum > target && endIndex - startIndex > 1) recursiveTwoSum(startIndex, endIndex - 1)
        else Array(startIndex + 1, endIndex + 1)
    }
    recursiveTwoSum(0, numbers.length - 1)
  }
}