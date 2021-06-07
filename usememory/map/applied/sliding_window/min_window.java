class Solution {
    public String minWindow(String s, String t) {
        // different approaches
        // brute force
        // try every starting location and every ending location O(n^2) and for each go over t O(n)
        
        // ideas
        // * can we use memory (maps & sets) to reduce computation
        // * two pointers for start and end?
        
        // reduce lookup of t = cannot do a set because can have duplicates - could do a map with key = character and v = number of times, or could sort the letters - but we might have extra letters in there so 
        // this comparison isn't great, nor is a deep comparison because need to search the entire array
        // reduce iterations of start and end -> could do start and end pointers 
        Map<Character, Integer> tLeftMap = new HashMap<>();
        // O(n)
        for(int i = 0; i < t.length(); i++){
            char c = t.charAt(i);
            if(tLeftMap.containsKey(c)){
                tLeftMap.put(c, tLeftMap.get(c) + 1);
            }
            else {
                tLeftMap.put(c, 1);
            }
        }
        // O(m)?
        int tLeftCount = t.length(); // want to avoid having to check values of tLeftMap each time
        int shortestSubstringStart = 0;
        int shortestSubstringFinish = s.length() + 1; // make longer than substring in case nothing
        int substringStart = 0;
        int substringFinish = 0;
        while(substringFinish < s.length()){
            char c = s.charAt(substringFinish);
            // c in t
            if(tLeftMap.containsKey(c)){               
                int cLeftCount = tLeftMap.get(c);
                // still c left to find
                if(cLeftCount > 0){
                    tLeftCount--;
                }
                // c left to find (can be negative) if have extra stored up
                tLeftMap.put(c, cLeftCount - 1);
                
                // still have t left
                // -> substringFinish++
                // all t found
                if(tLeftCount == 0){
                    if(substringFinish - substringStart < shortestSubstringFinish - shortestSubstringStart){
                            shortestSubstringFinish = substringFinish;
                            shortestSubstringStart = substringStart;
                    }
                    char cStart = s.charAt(substringStart);
                    // see if can substringStart++ and keep all t
                    // either because cStart not in t, or in t but we have extra
                    while(substringStart < substringFinish && 
                          (!tLeftMap.containsKey(cStart) || tLeftMap.get(cStart) < 0)){
                        
                        if(tLeftMap.containsKey(cStart)){
                            tLeftMap.put(cStart, tLeftMap.get(cStart) + 1);
                        }
                        cStart = s.charAt(++substringStart);
                        if(substringFinish - substringStart < shortestSubstringFinish - shortestSubstringStart){
                            shortestSubstringFinish = substringFinish;
                            shortestSubstringStart = substringStart;
                        }

                    }
                }
            }
            // c not in t

            // either way
            substringFinish++;
        }

        if(shortestSubstringFinish == s.length() + 1){
            return "";
        }
        else{
           return s.substring(shortestSubstringStart, shortestSubstringFinish + 1);
        }
    }
}