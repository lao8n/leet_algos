object Solution {
    def numIslands(grid: Array[Array[Char]]): Int = {
        // brute force
        // for every '1' bst/dst over all the 1s that are connected -> adding it to seen
        // and increment by 1
        // if do two for loops over every index + bst then you have O(2mn)
        // perhaps can do better where you use alternating bst across the whole matrix. so you only visit 
        // every point once
        
        // dynamic programming
        // way to use memory?
        // you could have a list of islands - and when two values are connected add them to the same list
        // perhaps you do one pass and they are all separate islands - and then you check for connections
        // and slowly merge the lists - will be very expesnive in terms of complexity
        // this is basically the union-find approach but instead of lists you use trees

        // graph approaches
        // dfs or bfs (need queues)
        // OPTIMIZATION = do not need separate seen matrix - can just set grid to 0

        // greedy approach
        // graph algorithm approach for connected components? perhaps counting the number of edges & vertices?
        // connected components is more relevant for minimal spanning trees and detecting cycles
        
        // seen is a grid of false which update to false
        val seen = Array.ofDim[Boolean](grid.length, grid(0).length)
        var islandCount = 0
        for (i <- 0 until grid.length){
            for(j <- 0 until grid(0).length){
                if(seen(i)(j) == false){
                    if(grid(i)(j) == '1'){
                        islandCount += 1
                    }
                    recursiveConnectedPoints((i, j), grid(i)(j))
                }
            }
        }
        def recursiveConnectedPoints(rc : (Int, Int), island : Char): Unit = {
            rc match {
                // base case - arrive at a point which is a different type or already seen
                case (r, c) if grid(r)(c) != island || seen(r)(c) == true =>
                // recursive case
                case (r, c) =>
                    seen(r)(c) = true
                    // no need for stack
                    getNeighbors(rc).foreach{
                        case nrc => recursiveConnectedPoints(nrc, island) 
                    }
            }
        }
        
        def getNeighbors(rc : (Int, Int)): List[(Int, Int)] = {
            rc match {
                case (r, c) => List((r - 1, c), (r + 1, c), (r, c - 1), (r, c + 1)).filter{
                    // remove invalid (off-grid) neighbors
                    // can't just do right and down because could have a weirdly connected shape
                    case (nr, nc) => nr < grid.length && nr >= 0 && nc < grid(0).length && nc >= 0 
                }
            }
        }
        islandCount
    }
}