object Solution {
  def trap(height: Array[Int]): Int = {
    // brute force approach 
    // for every square loop from left and right to find the maximum heights either side 
    // then take minimum for height
    
    // dynamic programming
    // only do the loop from left and right once each
    val leftMax = height.zipWithIndex
                        .foldLeft(Array.ofDim[Int](height.length), 0){
        case ((leftMaxs, currentMax), (h, i)) if h > currentMax => 
            leftMaxs(i) = h 
            (leftMaxs, h)
        case ((leftMaxs, currentMax), (h, i)) => 
            leftMaxs(i) = currentMax
            (leftMaxs, currentMax)
    }._1
    val rightMax = height.zipWithIndex
                          .foldRight(Array.ofDim[Int](height.length), 0){
        case ((h, i), (rightMaxs, currentMax)) if h > currentMax => 
            rightMaxs(i) = h 
            (rightMaxs, h)
        case ((h, i), (rightMaxs, currentMax)) => 
            rightMaxs(i) = currentMax
            (rightMaxs, currentMax)
    }._1
    height.zipWithIndex
          .foldLeft(0){
              case (total, (h, i)) => (total + Math.min(leftMax(i), rightMax(i)) - h)
          }
  }
}