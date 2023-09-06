object Insertion {
  def insertionSort(xs: List[Int]): List[Int] = {
    if (xs.isEmpty) Nil
    else insertElement(xs.head, insertionSort(xs.tail))
  }

  def insertElement(x: Int, xs: List[Int]): List[Int] = {
    if(xs.isEmpty || x <= xs.head) x :: xs
    else xs.head :: insertElement(x, xs.tail)
  }
}