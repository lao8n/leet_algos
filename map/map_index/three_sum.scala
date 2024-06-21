object Solution {
    // approach = we use dynamic programming to break the problem into a smaller problem
    //      three sum problem = two sum problem + one sum problem
    
    // three approaches to two sum problem 
    //  1. brute force with 2 for loops O(n^2) 
    //  2. use memory to save time - O(n) a HashMap (one pass) - we choose this approach
    //  3. sort the list and use two pointers O(n logn) because of sorting cost
    
    // choice between mutable and immutable data structure - choose the former as latter requires 
    // tail recursion on recursive calls - foldLeft is cleaner
    import scala.collection.mutable.Set
    import scala.collection.mutable.HashMap
    

    def threeSum(nums: Array[Int]): List[List[Int]] = {  
        // we build hashmap with one pass which we can re-use 
        // (rather than building at the same time as with two sum)
        val complementIndexMap = nums.zipWithIndex.foldLeft(Map[Int, Int]()){
            case (numsMap, (num: Int, numIndex: Int)) => numsMap + (num -> numIndex)    
        }
        // cannot use standard two sum approach because need to return all solutions 
        // therefore need to have additional solutions set which we add to as we find each solution
        // (rather than just returning immediately when we find the first solution)
        def twoSum(target: Int, targetIndex: Int, solutions: Set[List[Int]]): Set[List[Int]] = {
            nums.zipWithIndex
                .foldLeft(solutions){
                case (solutionsAcc, (num, numIndex)) => 
                        val complement = target - num
                        complementIndexMap.get(complement) match {
                            case Some(complementIndex) => 
                                if(List(targetIndex, numIndex, complementIndex).distinct.length == 3){
                                    solutionsAcc.add(List(-target, num, complement).sorted)
                                }
                                solutionsAcc
                            case None => solutionsAcc
                        }
            }
                
        }
        nums.zipWithIndex
            .distinct // getting memory limit exceeded when there are lots of repetitions
            .foldLeft(Set[List[Int]]()){
                case (solutions, (target, targetIndex)) => 
                    twoSum(-target, targetIndex, solutions)
            }.toList
    } 
}  