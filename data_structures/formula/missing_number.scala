object Solution {
    def missingNumber(nums: Array[Int]): Int = {
        // greedy or dynamic programming? 
        
        // brute force approach
        // loop over numbers - but not sure what one could check for - perhaps hypothesised numbers
        
        // sorting approach
        // O(n logn) to sort then O(n) to go over sorted numbers to see if there is a gap
        
        // rolling max-min - greedy
        // you store rolling max and min for O(1) cost and loop over numbers O(n) runtime. 
        // if max - min + 1 == nums.length then all numbers in the range and missing is max + 1
        // else take min + 1 and loop over numbers incrementing everytime you find x + 1 - doesn't work as could miss item
        
        // dynamic programming approach
        // if have max - min + 1 == num.length then missing is max + 1
        // so given set of numbers recursively find max and min of remaining numbers until condition hold
        // [3,0,1]
        // i = 0, j = 2, max_left (3, 1) = 3 - 1, min_left (3, 1) = 1 + 1, therefore suggested range left is 2-2,
        // i = 1, j = 1, max_left = 2, min_left - unchanged because found outside range
        
        // [9,6,4,2,3,5,7,0,1]
        // i = 0, j = 8, max_left(9, 1) = 9 - 1 = 8, min_left (9, 1) = 1 + 1 = 2, suggested range = 2-8
        // i = 1, j = 7, max_left(8, 6, 0) = 8, min_left(2, 6, 0) = 0, suggested range 1-8 - but shouldn't be 1 because already had it 
        
        // [9,6,4,2,3,5,7,0,1]
        // [0, 1] = 2
        // [7, 0, 1] = invalid
        // [5, 7, 0, 1] = 
        
        // [9,6,4,2,3,5,7,0,1]
        // [9, 6] = min_missing = 7, min=6, max_missing = 8, max=9 
        // [4] = min = 4, min_missing=5, max=9, max_missing =8
        // [2] = min = 2, min_missing=3, max=9, max_missing =8
        // [3] = min = 2, min_missing=4 (forgot) max = 9, max_missing=8
        // [5] = min = 2, min_missing=4 max = 9, max_missing = 8
        // [7] = min = 3 min_missing = 4
        
        // [9,6,4,2,3,5,7,0,1]
        // [9, 6] min_missing_range = 7-8, max_missing_range = 7-8
        // [4] min_missing_range = 5-(7-2), max_missing_range = 7-8
        // [2] min_missing_range = 3-(5-2), max_missing_range = 7-8
        // [3] min_missing_range = na, max_missing_range = 7-8
        
        // key problem i'm running into is 
        // 1. forgetting that i've already got say 4 and then having all the numbers below four filled in but
        // thinking that 4 is the new missing_minimum
        
        // suppose you have a list of numbers 
        // [0, 1, 2, 3, 4, 5, 6, 7, 8, 9]
        // if you choose two numbers you partition into either 1 range of missing 1-8
        // or 3 ranges which is in between , below and above
        // then if you choose another element you either keep 3 ranges or get 4 ranges or maybe go back to 2 ranges
        // worst case scenario you have n/2 ranges which you can fit in a array of n
        // [9,6,4,2,3,5,7,0,1]
        // [9, 6] => max = 6 + (n - 1) = 5 + n, min = 9 - (n - 1) = 8 - n
        // [-1,...,14] // if -1 then -1 has to exist therefore missing is [0-5, 7-8, 10-14] 
        // [4] = > max = 4 + (n - 1) = 3 + n = 12, 
        // [0-3, 5, 7-8, 10-14]
        // [2] => max = 2 + (n - 1) = 10
        // [0-1, 3-3, 5, 7-8, 10-10]
        // [3]
        // [0-1, 5, 7-8, 10]
        // [5]
        // [0-1, 7-8, 10]
        // [7]
        // [0-1, 8, 10]
        // [0]
        // [1, 8]
        // [1]
        // [8]
        
        // half sorting algorithm - keep tracking of min and max
        // [9,6,4,2,3,5,7,0,1]
        // [1,6,4,2,3,5,7,0,9] min = 1@0, max = 9@8
        
        // guess at what solution is - some sort of variation on boyer-moore majority vote algorithm
        // insted of just actual min and max that gotten, use the fact that have everyone element gotten
        // so we must know the complete min and max range - then every element we get -if outside - reduces the range by at least 1 or if inside it 
        
        // looking at the solution 
        // i got the sorting approach - i can understandt he hashset approach
        // key thing i missed is the greedy approach with the summation - i don't understand the xor approach
        // use gauss' formula for sum of numbers
        nums.length*(nums.length+1)/2 - nums.sum
    }
}