object Solution {
    def isPalindrome(s: String): Boolean = {
        // brute force approach 
        // * two pointers and compare each one until get to the middle - O(n) 
        // * use a stack to increment and decrement to O(n) - but uses memory - and difficult to know if you 
        // are half-way
        
        // q: foldLeft vs recursion?
        
        // situations need to handle
        // * TODO: if nonalphanumeric then move on
        // * TODO: upper and lower case
        // * TODO: even and odd length strings
        def isPalindromeRecursive(leftIndex: Int, rightIndex: Int): Boolean = {
            // base case
            // scenario 1: [leftIndex, x, rightIndex] => [, leftIndex=rightIndex, ] 
            // scenario 2: [leftIndex, rightIndex] => [rightIndex, leftIndex]
            if(leftIndex >= rightIndex) true
            else {
                // recursive cases
                if(!s(leftIndex).isLetterOrDigit) isPalindromeRecursive(leftIndex + 1, rightIndex)
                else if(!s(rightIndex).isLetterOrDigit) isPalindromeRecursive(leftIndex, rightIndex - 1)
                else if(s(leftIndex).toLower == s(rightIndex).toLower)isPalindromeRecursive(leftIndex + 1, rightIndex - 1)
                else false
            }
        } 
        isPalindromeRecursive(0, s.length - 1)
    }
}