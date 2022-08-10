object Solution {
    def combinationSum(candidates: Array[Int], target: Int): List[List[Int]] = {
        // brute force
        // given a target
        // for loop over all the different combinations and then see if it equals the target
        // target^n          

        // not two sum variation (hashmap or sort + two-pointers) because unlimited use
        
        // dynamic programming 
        // target = 7 = ans = []
        // target = 3 = ans = [4, ]
        val listCandidates = candidates.toList
        // 1. how are we going to combine the elements ? accumulator, returning value or maybe mutable? don't 
        //    care about order so probably okay
        // 2. how do we get uniqueness? cheat by set with list afterwards 
        //    better solution is you need to add a value that is larger than values so far 
        // 3. unclear on whether we can append options -> double recursion - independent 
        // 4. how do we dump those that don't work - just return List not List(List())
        
        def recursiveCombinationSum(target: Int, last : Int) : List[List[Int]] = {
            // base case
            listCandidates.filter(_ <= Math.min(target, last)) match {
                case Nil if target == 0 => List(List())
                case Nil if target > 0 => List()
                case filteredCandidates => filteredCandidates.flatMap{c =>
                    recursiveCombinationSum(target - c, c).map(ans => c :: ans)
                }
            }
        }
        recursiveCombinationSum(target, target)
    }
}