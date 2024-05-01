final case class Tree[+A](value: A, children: List[Tree[A]])

object Solution {
    // DFS by right.
    def find[A](value: A)(tree: Tree[A]): Option[Tree[A]] =
    LazyList.unfold(List(tree)) {
        case Nil => None
        case tree :: remaining => Some((tree, tree.children reverse_::: remaining))
    }.find(tree => tree.value == value)

    // DFS by left.
    def find[A](value: A)(tree: Tree[A]): Option[Tree[A]] =
    LazyList.unfold(List(tree)) {
        case Nil => None
        case tree :: remaining => Some((tree, tree.children ::: remaining))
    }.find(tree => tree.value == value)
}