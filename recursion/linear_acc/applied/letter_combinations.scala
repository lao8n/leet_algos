object Solution {
    def letterCombinations(digits: String): List[String] = {
        // brute force
        // maintain a map for what each letter stands for 
        // iterating over each combination - given that each number represents 3 or 4 letters we have
        // 4 ^ n and 
        val phoneMap = Map(('2' -> "abc"), ('3' -> "def"), ('4' -> "ghi"), ('5' -> "jkl"), ('6' -> "mno"), 
                           ('7' -> "pqrs"), ('8' -> "tuv"), ('9' -> "wxyz"))
        // greedy problem where you can take each element and combine - independently of the other steps
        // can't just recurse from the end because it will be in the wrong order
        def recursiveLetterCombinations(digits: String, acc : List[String]): List[String] = {
            digits match {
                // base case
                case "" => acc 
                // recursive case
                case digits => 
                    recursiveLetterCombinations(digits.tail, 
                                               // ["ad", "ae"...etc] -> "ad" + "a" or  "b" or "c"
                                               acc.flatMap(ans => phoneMap(digits.head).map(elem => ans + elem)))
            }
        }
        if(digits == "") 
            List()
        else
            recursiveLetterCombinations(digits, List(""))
        
    }
}