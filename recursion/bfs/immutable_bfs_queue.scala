object Solution {
  final case class Tree[+A](value: A, left: Option[Tree[A]], right: Option[Tree[A]])

  def find[A](value: A)(tree: Tree[A]): Option[Tree[A]] = {
    import scala.collection.immutable.Queue

    @annotation.tailrec
    def bfs(queue: Queue[Tree[A]]): Option[Tree[A]] =
        queue.dequeueOption match {
        case None => None

        case Some((tree, remaining)) => tree match {
            case Tree(`value`, _, _) => Some(tree)
            case Tree(_, Some(left), Some(right)) => bfs(queue = remaining.enqueue(left).enqueue(right))
            case Tree(_, Some(left), None) => bfs(queue = remaining.enqueue(left))
            case Tree(_, None, Some(right)) => bfs(queue = remaining.enqueue(right))
            case Tree(_, None, None) => bfs(queue = remaining)
        }
        }

    bfs(queue = Queue(tree))
  }
}