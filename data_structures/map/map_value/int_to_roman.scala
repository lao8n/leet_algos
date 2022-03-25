import static java.util.Map.entry;    
class Solution {
    public String intToRoman(int num) {
        Map<Integer, Character> intToRomanMap = Map.of(
            1000, 'M',
            500, 'D',
            100, 'C',
            50, 'L',
            10, 'X',
            5, 'V',
            1, 'I'
        );
        StringBuffer romanNumeralString = new StringBuffer();
        int tenOrFive = 10;
        int i = 1000;
        while(i >= 1){
            num = calculateRomanNumeral(tenOrFive, romanNumeralString, intToRomanMap, num, i);
            if(tenOrFive == 10){
                tenOrFive = 5;
                i = i / 2;
            } 
            else {
                tenOrFive = 10;
                i = i / 5;
            }
        }
        return romanNumeralString.toString();
    }
    
    private int calculateRomanNumeral(
        int tenOrFive,
        StringBuffer romanNumeralString, 
        Map<Integer, Character> intToRomanMap, 
        int num, 
        int multipleOfTenOrFive){
        // 1000, 100, 10 or 500, 50 or 5
        while(num >= multipleOfTenOrFive){
            romanNumeralString.append(intToRomanMap.get(multipleOfTenOrFive));
            num -= multipleOfTenOrFive;          
        }
        // 900, 90, 9 or 400, 40 or 4
        if(num >= multipleOfTenOrFive - multipleOfTenOrFive / tenOrFive){
            romanNumeralString.append(intToRomanMap.get(multipleOfTenOrFive / tenOrFive));
            romanNumeralString.append(intToRomanMap.get(multipleOfTenOrFive));
            num -= multipleOfTenOrFive - multipleOfTenOrFive / tenOrFive;
        }
        return num;
    }   
}