object Solution {
    def sortedSquares(nums: Array[Int]): Array[Int] = {
        // brute force
        // loop through squaring which costs O(n)
        // can then sort with O(n log n)
        
        // for loop until find last negative number and first positive number O(n)
        // then square and add smallest
        // kind of similar to quick sort where we partition around a pivot - here the pivot is 0
        val (lessThanZero, greaterThanZero) = nums.partition(_ < 0)
        def sortedSquareRecursive(lessThanZero : Array[Int], greaterThanZero : Array[Int]): Array[Int] = {
            (lessThanZero, greaterThanZero) match {
                // base cases
                case (Array(), gt) => gt 
                case (lt, Array()) => lt.reverse
                // recursive case
                case (lt, gt) if lt.last <= gt.head =>
                    lt.last +: sortedSquareRecursive(lt.dropRight(1), gt)
                case (lt, gt) if gt.head < lt.last => 
                    gt.head +: sortedSquareRecursive(lt, gt.tail)
            }
        }
        sortedSquareRecursive(lessThanZero.map(x => x * x), greaterThanZero.map(x => x * x))
    }
}