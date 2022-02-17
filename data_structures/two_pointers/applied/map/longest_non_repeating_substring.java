class Solution {
    public int lengthOfLongestSubstring(String s) {
        // naive solution
        // for loop of all possible starting positions and 
        // for loop of all possible finishing positions
        // for o(n^2) complexity
        
        // maintain a rolling set of values in the substring.
        // i.e. start_index = 0, end_index = 0
        // loop through whilst finding chars not in set
        // if already in set then jump start from where that value was in set
        // therefore should use a map to know where to jump to
        Map<Character, Integer> substringMap = new HashMap<>();
        
        int substringStart = 0;
        int substringFinish = 0;
        int longestSubstring = 0;
        while(substringFinish < s.length()){
            char c = s.charAt(substringFinish);
            // c already exists so we jump substringStart to where that last 
            // instance was
            if(substringMap.containsKey(c)){
                int lastCIndex = substringMap.get(c);
                while(substringStart <= lastCIndex){
                    char k = s.charAt(substringStart);
                    substringMap.remove(k);
                    substringStart++;
                }
            }
            else {
                // can only possibly be longer if didn't have to remove from map
                if(substringFinish - substringStart >= longestSubstring){
                    longestSubstring = substringFinish - substringStart + 1;
                }
            }
            substringMap.put(c, substringFinish);
            substringFinish++;
        }
        return longestSubstring;
    }
}

public Solution{
    public int lengthOfLongestSubstring(String s){
        int longestSubstring = 0;
        Map<Character, Integer> substringMap = new HashMap<>();

        for(int substringStart = 0, substringFinish = 0; substringFinish < s.length(); substringFinish++){
            char c = s.charAt(substringFinish);
            if(substringMap.contains(c)){
                substringStart = Math.max(substringMap.get(c), substringStart);
            }
            longestSubstring = Math.max(longestSubstring, substringFinish - substringStart + 1);
            substringMap.put(c, substringFinish + 1);
        }
        return longestSubstring;
    }
}