object Solution {
    def trap(height: Array[Int]): Int = {
        // brute force approach 
        // for every square - test left and right to see if is surrounded
        
        // memory -> stack
        // stack of indices is a rolling list of smallest bars
        import scala.collection.mutable
        def recursiveTrap(hI :Int, stackLastMinHeights : mutable.Stack[Int], total : Int): Int = {
            // base case
            if(hI == height.length) return total
            // iterative base case
            else if(stackLastMinHeights.isEmpty){
                recursiveTrap(hI + 1, stackLastMinHeights.push(hI), total)
            }
            // if current height more add trapped water
            else if(height(hI) > height(stackLastMinHeights.top) && stackLastMinHeights.length == 1){
                stackLastMinHeights.pop
                recursiveTrap(hI, stackLastMinHeights, total)
            }
            else if(height(hI) > height(stackLastMinHeights.top) && stackLastMinHeights.length > 1){
                val lastMinHeightI = stackLastMinHeights.pop
                val beforeLastMinHeightI = stackLastMinHeights.top
                val widthTrapped = hI - beforeLastMinHeightI - 1
                val heightTrapped = Math.min(height(hI), height(beforeLastMinHeightI)) - height(lastMinHeightI)
                recursiveTrap(hI, stackLastMinHeights, total + widthTrapped * heightTrapped)
            }
            else {
                // if current height smaller push to stack
                recursiveTrap(hI + 1, stackLastMinHeights.push(hI), total)
            }
        }
        recursiveTrap(0, mutable.Stack[Int](), 0)
    }
}