/**
 * Definition for a Node.
 * class Node(var _value: Int) {
 *   var value: Int = _value
 *   var neighbors: List[Node] = List()
 * }
 */

object Solution {
    import scala.collection.mutable
    def cloneGraph(graph: Node): Node = {
        val visitedNodes = mutable.HashMap[Int, Node]()
        def recursiveCloneGraph(graph: Node): Node = {
            graph match {
                case null => null
                case graph if visitedNodes.contains(graph.value) => visitedNodes.getOrElse(graph.value, new Node(0))
                case graph => {
                    val cloneNode = new Node(graph.value)
                    visitedNodes += (graph.value -> cloneNode)
                    // no need for stack - just immediately cycle through
                    graph.neighbors.foreach {
                        n => cloneNode.neighbors = recursiveCloneGraph(n) :: cloneNode.neighbors
                    }
                    cloneNode
                }
            }
        }
        recursiveCloneGraph(graph)
    }
}