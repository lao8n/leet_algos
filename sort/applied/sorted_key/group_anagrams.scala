object Solution {
    def groupAnagrams(strs: Array[String]): List[List[String]] = {
        // greedy or dynamic programming approach
        // greedy approach where in a single pass you build up a map of letters key -> list index
        // letters key can be done in two ways 
        // 1. sort each word 
        // 2. multiply each word by a number to give unique output
        // problem with comparing on count is you then need to do a deep comparison which is difficult
        strs.toList.groupBy(_.sorted).values.toList

        // rather than automatically folding left over a collection look to group by
        // if possible
    }
}