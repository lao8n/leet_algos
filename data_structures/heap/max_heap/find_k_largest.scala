object Solution {
  def findKthLargest(nums: Array[Int], k: Int): Int = {
    val queue = scala.collection.mutable.PriorityQueue.empty[Int]
    nums.foreach { n =>
      queue.enqueue(-n)
      if (queue.length > k) queue.dequeue()
    }
    -queue.head
  }
}