1. the brute force approach loops over all start and end indices and then 
loops over the substring to test if it is a set giving O(n^3) complexity
2. a sliding window with a set means we can remove one of the O(n) by 
checking the existence in the substring in O(1) time. Then we can save
another O(n) by only checking the next end index not all of them for each 
starting index. You still have O(2n) complexity however as you try each 
starting index and (potentially) every ending index
3. replace the set with a map of value to index so you know which index to 
jump to

improvements
1. you can further improve by not removing keys as you increment but overriding
instead
2. use max to figure out which index to use.

differences across languages
essentially the same except for immutable data structures like scala with 
recursion