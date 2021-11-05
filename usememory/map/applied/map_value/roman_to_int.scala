class Solution {
    public int romanToInt(String s) {
        Map<Character, Integer> romanToIntMap = Map.of(
            'M', 1000,
            'D', 500,
            'C', 100,
            'L', 50,
            'X', 10,
            'V', 5,
            'I', 1
        );
        int num = 0;
        for(int i = 0; i < s.length(); i++){
            int j = i + 1;
            int cI = romanToIntMap.get(s.charAt(i));
            if(j < s.length() && cI < romanToIntMap.get(s.charAt(j))){
                num -= cI;
            } else {
                num += cI;
            }
        }
        return num;
    }
}