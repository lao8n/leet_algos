object Solution {
    def fourSum(nums: Array[Int], target: Int): List[List[Int]] = {
        // greedy or dynamic programming? probably latter because can do two-sum approach ++
        
        // two sum has 3 approaches
        // 1. brute force 2x for loops
        // 2. hashmap
        // 3. sort + two-pointers
        
        // four-sum
        // 1. 2 for loops + hashmap - O(n^3)
        // 2. 2 for loops + two-pointers - ~O(n^3)
        // 3. other approaches? build a map of (sum -> (i, j)) to indices O(n^2) - then iterate through keys (i.e. sums) to see if (target - sum exists)
        
        import scala.collection.mutable.HashMap
        import scala.collection.mutable.Set
        var pairMap = HashMap[Int, List[(Int, Int)]]()
        for(i <- 0 until nums.size){
            for(j <- i + 1 until nums.size){
                // problem is if have two sums the same we only keep one and we need to find all the solutions
                val pairSum = nums(i) + nums(j)
                pairMap = pairMap += (pairSum -> ((i,j) :: pairMap.getOrElse(pairSum,List[(Int, Int)]())))
            }
        }
        pairMap.keySet
             .foldLeft(Set[List[Int]]()){
                 (solutions, pairSum) => 
                    pairMap.get(target - pairSum)
                            // need to check for duplicate indices
                           .map{listComplementIndices => listComplementIndices
                                .foreach{ case (complementIndex1: Int, complementIndex2: Int) => 
                                    val listIndices = pairMap.getOrElse(pairSum, List[(Int, Int)]())
                                    listIndices.foreach {
                                        case (index1, index2) =>
                                            if(List(index1, index2, complementIndex1, complementIndex2).distinct.length == 4)
                                                solutions.addOne(List(nums(index1), nums(index2), nums(complementIndex1), nums(complementIndex2)).sorted) 
                                    }
                                }
                                solutions
                                    
                            }.getOrElse(solutions)

             }.toList
    }
}