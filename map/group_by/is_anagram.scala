object Solution {
    def isAnagram(s: String, t: String): Boolean = {
        // we need a way to test uniqueness
        // 1. sorted
        // s.sorted == t.sorted
        // 2. calculate hash with prime numbers 
        // s.groupBy(identity).mapValues(_.length) == t.groupBy(identity).mapValues(_.length)
        val sCounts = s.groupBy(identity).mapValues(_.length)
        val tCounts = t.groupBy(identity).mapValues(_.length)
        if(sCounts.keySet != tCounts.keySet) false
        else sCounts.keySet.filter{k => sCounts(k) != tCounts(k)}.size == 0
    }
}