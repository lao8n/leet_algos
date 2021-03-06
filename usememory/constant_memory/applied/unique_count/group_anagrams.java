class Solution {
    public List<List<String>> groupAnagrams(String[] strs) {
        // brute force approach is you take each element and see if it fits in one of the buckets
        // each bucket is determined by the characters used (to see if bucket matches we can sort the string)
        // to give it a label we use a map with key being sorted string and value being a list of items
        // however this approach requires building the list again separately so instead we just need a map with
        // an index saying which sublist to add the item to 
        Map<String, Integer> charactersToList = new HashMap<>();
        List<List<String>> result = new ArrayList<>();
        for(int i = 0; i < strs.length; i++){
            String anagramKey = sortString(strs[i]);
            if(charactersToList.containsKey(anagramKey)){
                Integer listIndex = charactersToList.get(anagramKey);
                List<String> listToUpdate = result.get(listIndex);
                listToUpdate.add(strs[i]);
                result.set(listIndex, listToUpdate);
            }
            else{
                charactersToList.put(anagramKey, result.size());
                result.add(new ArrayList<String>(Arrays.asList(strs[i])));
            }
        }
        return result;
    }
    
    public String sortString(String inputString)
    {
        char tempArray[] = inputString.toCharArray();
        Arrays.sort(tempArray);
        return new String(tempArray);
    }

    // can actually do better where directly use map to store values and then just output values at the end
    class Solution {
    public List<List<String>> groupAnagrams(String[] strs) {
        if (strs.length == 0) return new ArrayList();
        Map<String, List> ans = new HashMap<String, List>();
        for (String s : strs) {
            char[] ca = s.toCharArray();
            Arrays.sort(ca);
            String key = String.valueOf(ca);
            if (!ans.containsKey(key)) ans.put(key, new ArrayList());
            ans.get(key).add(s);
        }
        return new ArrayList(ans.values());
    }
}
}