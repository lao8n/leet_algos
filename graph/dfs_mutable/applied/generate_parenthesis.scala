import scala.collection.mutable.ListBuffer

object Solution {
  def generateParenthesis(n: Int): List[String] = {
    val acc = ListBuffer[String]()
    def dfs(open: Int, closed: Int, curr: String): Unit = {
      if (open == n) {
        val combination = curr + (")" *  (n - closed ))
        acc.addOne(combination)
      }
      else if (closed > open) ()
      else {
        dfs(open + 1, closed, curr + "(")
        dfs(open, closed + 1, curr + ")")
      }
    }
    dfs(1, 0, "(")
    acc.toList
  }
}