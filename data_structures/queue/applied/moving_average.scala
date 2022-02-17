class MovingAverage(_size: Int) {

    /** Initialize your data structure here. */
    
    // brute force
    // maintain a list of all values and re-calculate it every time
    
    // dynamic approach
    // memoization argument where same calculations 
    // for example if you have [n-3, n-2, n-1, n, n+1]
    // if the window size is 4 at point n you have (n-3 + n-2 + n-1 + n)/4
    // then for the next value along you have (n-2 + n-2 + n + n+1) / 4
    // so of that computation a lot is shared where the difference is -(n-3)/4 + (n+1)/4
    // however you still need to keep a list of all the previous values to know what to minus
    // i.e. it costs you O(1) in processing time - you just take the first value but it costs you O(n) in memory
    
    import scala.collection.mutable
    // can we use immutable data structures and recursion?
    val q = mutable.Queue[Int]()
    var sum = 0D
    def next(`val`: Int): Double = {
        // dequeue 
        if(q.size >= _size) {
            sum -= q.dequeue()
        }
        sum += `val`
        q.enqueue(`val`)
        sum / q.size
    }

}

/**
 * Your MovingAverage object will be instantiated and called as such:
 * var obj = new MovingAverage(size)
 * var param_1 = obj.next(`val`)
 */