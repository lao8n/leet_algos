object Solution {
    def multiply(mat1: Array[Array[Int]], mat2: Array[Array[Int]]): Array[Array[Int]] = {
        // brute force
        // for loop over every row & column combination for O(m x k x k x n)
        
        // use memory (dynamic programming)
        // first matrix only uses rows, and second matrix only uses columns
        // we re-use each row in the first matrix n and we re-use each column in the second matrix m times
        // so we could have a map with (0, 0) -> 1, (0, 1) -> 0, (0, 2) -> 0 etc.
        // which means complexity of O(m x k + k x n)
        // however, we can do even better by only storing in the map non-zero values (we are told it is sparse)
        // and then searching for keys
        val mat1Map = mat1.zipWithIndex
                          .flatMap{ case (row, rowIndex) => 
                                row.zipWithIndex.map{case (v,colIndex) => (rowIndex, colIndex, v)}}
                          .filter{ case (_, _, v) => v != 0}
                          .groupBy{ case (x, y, v) => x }
        val mat2Map = mat2.zipWithIndex
                          .flatMap{ case (row, rowIndex) => 
                                row.zipWithIndex.map{case (v,colIndex) => (rowIndex, colIndex, v)}}
                          .filter{ case (_, _, v) => v != 0}
                          .groupBy{ case (x, y, v) => x }
        // could potentially optimise that you pick the smaller map to foldleft over 
        mat1Map.foldLeft(Array.fill[Int](mat1.length, mat2(0).length)(0)){
            case (solution, (r1, row1)) => 
                row1.foreach{
                    case (r1 : Int, c1 : Int, v1 : Int) => 
                        mat2Map.getOrElse(c1, Array[Int]()) match {
                            case row2 => row2.foreach{
                                case (r2 : Int, c2 : Int, v2 : Int) => solution(r1)(c2) = (solution(r1)(c2) + v1 * v2)
                            }
                        }
                }
                solution 
        }
    }
}