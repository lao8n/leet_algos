object Solution {
  def reverse[T](xs: List[T]): List[T] = xs match {
      case List() => List()
      case x :: xs1 => reverse(xs1) ::: List(x)
  }
}