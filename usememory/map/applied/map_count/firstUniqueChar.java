class Solution {
    public int firstUniqChar(String s) {
        // brute force approach
        // loop through considering each character for O(n)
        // and for each character look for duplicates another O(n)
        // for O(n^2)
        
        // memory?
        // for each character we consider we look for duplicates, this look up can be done with 
        // O(1) time 
        // if we did a map with character and number of occurences we would have to do O(2n)
        // because one loop to add to values another loop to see which is first
        // perhaps something clever can be done with an ordered map or storing something different.
        // or perhaps even doing it in one loop 
        // perhaps the answer is to use a linkedhashmap to avoid the second O(n) loop which preserves
        // insertion order
        
        // sorting?
        // sorting costs O(n logn) so is already worst then a O(2n) memory approach
        // the adv of sorting is then you can identify the set of non duplicates quite quickly, 
        // but that would cost another O(n) memory 
        // and then you can go through list again to see which one is first
        // but actually no point because do not have index
        
        // dynamic?
        // can we ask the same question but perhaps with a smaller subset of the data?
        // yes but for 'first' type questions doesn't work that well because have to maintain all candidates?
        
        // greedy? 
        // for greedy we typically want to do some sort of rolling sum, count, max, min etc. 
        // can't thikn of anything obvious we could try
        
        Map<Character, Integer> characterCount = new HashMap<>();
        for(int i = 0; i < s.length(); i++){
            if(characterCount.containsKey(s.charAt(i))){
                characterCount.put(s.charAt(i), characterCount.get(s.charAt(i)) + 1);
            }
            else {
                characterCount.put(s.charAt(i), 1);
            }
        }
        for(int i = 0; i < s.length(); i++){
            if(characterCount.get(s.charAt(i)) == 1){
                return i;
            }
        }
        return -1;
    }
}