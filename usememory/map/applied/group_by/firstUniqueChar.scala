object Solution {
    def firstUniqChar(s: String): Int = {
        // brute force approach
        // for loop over values O(n)
        // for loop checking for duplicates
        // but O(n^2)
        
        // sorting?
        
        // memory?
        // single loop through adding to a hashmap
        // then loop through again looking for duplicates, can probably do something cleaner with a filter
        // etc.
        
        // greedy? 
        
        // dynamic?
        
        // s.foldLeft(HashMap[Character, Integer]()){
        //     case (char, charCount) => charCount + (char, charCount.getOrElse(char, 0) + 1)
        // }
        s.zipWithIndex
         .groupBy(_._1) // returns hashmap (c -> ((c, index1), (c, index2)))
         .filter(_._2.length == 1) // look for unique i.e. only those hashmap(c -> (c, index)) 
         .flatMap(_._2) // HashMap((t, 3), (l, 0), etc.)
         .minByOption(_._2) // (l, 0)
         .map(_._2) // Some(0) or None
         .getOrElse(-1) // 0 or -1
    }
}