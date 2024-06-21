object Solution {
    def generateParenthesis(n: Int): List[String] = {
        // brute force
        // generate all possible sequences of '(' and ')' and check for validity with a stack
        
        // use memory - stack? 
        
        // greedy
        
        // dynamic programming 
        // n = 3 ==  n = 2
        // n = 2 "(())" "()()"
        // iterative step 
        // before "() + (())", "() + ()()" + 
        // after "(()) + ()" + "()()() + ()" +
        // around "((()))" + "(()())"
        // this doesn't work because cannot handle     "(())(())"
        // def recursiveGenerateParenthesis(n: Int) : Set[String] = {
        //     n match {
        //         // base case
        //         case 1 => Set("()")
        //         // iterative case
        //         case n => {
        //             val nSubset = recursiveGenerateParenthesis(n - 1) // list of strings
        //             // backtracking
        //             nSubset.flatMap{
        //                 nSub => Set("()" + nSub, // before
        //                              "(" + nSub + ")",  // around
        //                              nSub + "()") // after
        //             }
        //         }
        //     }
        // }
        // recursiveGenerateParenthesis(n).toList
        
        n match {
            case 0 => List("")
            case n =>
                for {
                    m <- (0 to n - 1).toList
                    x <- generateParenthesis(m)
                    y <- generateParenthesis(n-1-m)
                } yield ("(" ++ x ++ ")" ++ y)
        }
    }
}
