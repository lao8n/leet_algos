object mergesort {
  def mergeSort(xs: List[Int]): List[Int] = {
    val leftPartition, rightPartition = xs.split(xs.length / 2)
    merge(mergeSort(leftPartition), mergeSort(rightPartition))
  }

  def merge(leftPartition: List[Int], rightPartition: List[Int]){
    (leftPartition, rightPartition) match {
      case(Nil, rightPartition) => rightPartition
      case(leftPartition, Nil) => leftPartition
      case(leftHead :: leftTail, rightHead :: rightTail) =>
        if(leftHead < rightHead) leftHead :: merge(leftTail, rightPartition)
        else rightHead :: merge(left, rightTail)
    }
  }
}