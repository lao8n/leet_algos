final case class Tree[+A](value: A, children: List[Tree[A]])

object Solution {
    // BFS
    def find[A](value: A)(tree: Tree[A]): Option[Tree[A]] =
    LazyList.unfold(Queue(tree)) { queue =>
        queue.dequeueOption.map {
        case (tree, remaining) => (tree, remaining.enqueueAll(tree.children))
        }
    }.find(tree => tree.value == value)
}