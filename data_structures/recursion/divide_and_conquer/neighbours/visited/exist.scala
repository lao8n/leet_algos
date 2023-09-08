object Solution {
  def exist(board: Array[Array[Char]], word: String): Boolean = {
    val maxRow = board.length - 1
    val maxCol = board.head.length - 1

    def inner(w: List[Char], r: Int, c: Int, used: Set[(Int, Int)]): Boolean =
      if (w.isEmpty) true
      else if (r < 0 || r > maxRow || c < 0 || c > maxCol || used.contains((r, c))) false
      else if (used.nonEmpty && board(r)(c) != w.head) false
      else if (board(r)(c) == w.head)
        inner(w.tail, r, c - 1, used + ((r, c))) || // left
        inner(w.tail, r - 1, c, used + ((r, c))) || // up
        inner(w.tail, r, c + 1, used + ((r, c))) || // right
        inner(w.tail, r + 1, c, used + ((r, c))) || // down
        used.isEmpty && inner(w, if (c == maxCol) r + 1 else r, if (c == maxCol) 0 else c + 1, used)
      else
        inner(w, if (c == maxCol) r + 1 else r, if (c == maxCol) 0 else c + 1, used)

    inner(word.toList, 0, 0, Set.empty[(Int, Int)])
  }
}