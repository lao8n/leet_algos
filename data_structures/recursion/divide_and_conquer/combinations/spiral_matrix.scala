object Solution {
    def spiralOrder(matrix: Array[Array[Int]]): List[Int] = {
        // greedy or the dynamic programming approach
        // could sequentially calculate each index but maybe can actually calculate in parallel?
        // key is there are different loops and we can determine the loop that it is in by how far 
        // from edge that we are
        // dynamic programming approach is how can we get the same problem over and over again
        // but on different data
        // you have each ring - but actually if you transpose and reverse the data it is the same
        // problem for every edge with one edge remove
        // couldn't see a clever formula for a greedy parallel approach
        def recursiveSpiralOrder(matrix: Array[Array[Int]]): List[Int] = matrix match {
            case matrix if matrix.isEmpty => List()
            case matrix if matrix.length == 1 => matrix.head.toList
            case matrix => matrix.head.toList ::: recursiveSpiralOrder(matrix.tail.transpose.reverse)
        }
        recursiveSpiralOrder(matrix)
    }
}
