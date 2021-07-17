final case class Tree[+A](value: A, left: Option[Tree[A]], right: Option[Tree[A]])

object Solution {

    def find[A](value: A)(tree: Tree[A]): Option[Tree[A]] = {
        @annotation.tailrec
        def dfs(stack: List[Tree[A]]): Option[Tree[A]] =
            stack match {
            case Nil => None

            case tree :: remaining => tree match {
                case Tree(`value`, _, _) => Some(tree)
                case Tree(_, Some(left), Some(right)) => dfs(stack = left :: right :: remaining)
                case Tree(_, Some(left), None) => dfs(stack = left :: remaining)
                case Tree(_, None, Some(right)) => dfs(stack = right :: remaining)
                case Tree(_, None, None) => dfs(stack = remaining)
            }
        }

        dfs(stack = List(tree))
    }
}