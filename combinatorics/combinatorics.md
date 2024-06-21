Combinations = two main approaches where either build path on
1. forward pass - incl. loop over options until reach the end, backtrack if shallow copies
2. backward pass - incl. loop over candidates and over solutions

Subsets = same as combinations except add each partial path to solutions 

Permutations = same as combinations except 
1. forward pass - loop over candidates with seen map
2. backward pass - loop over candidates and over solutions but in candidates accumulate forward different subsets removing that candidate