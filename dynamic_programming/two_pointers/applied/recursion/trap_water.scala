object Solution {
  def trap(height: Array[Int]): Int = {
    def recursiveTrap(l: Int, r: Int, level: Int, total: Int): Int = {
      if (l >= r) {
        total
      } else {
        (l, r) match {
          case (l, r) if (height(l) <= height(r)) => 
            recursiveTrap(l+1, r, Math.max(level, height(l)), total+Math.max(0, level-height(l)))
          case (l, r) if (height(l) > height(r)) => 
            recursiveTrap(l, r-1, Math.max(level, height(r)), total+Math.max(0, level-height(r)))
        }
      }
    }
    recursiveTrap(0, height.length-1, 0, 0)
  }
}