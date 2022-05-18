There are types of problems:
1. permutations n! = order does matter
2. combinations n C k = order doesn't matter
3. subsets 2^n

Core approaches
* Recurse forwards to [] base case, recurse backwards looping over solutions adding candidates
* Recurse forwards with accumulator building to combination as base case, recurse backwards appending solutions
* Recurse forwards with accumulator and pointer to solutions to combination as base case and appending solutions

Key difference between permutations and combinations
* Permutations you loop over candidates and solutions, with combinations you loop over just one

Permutations 
* Generate = loop over candidates
    1. Loop over list as arg with `[:i]` and `[i+1:]` `permute.go`
* Order does matter = how keep different permutations
    1. In each recursive call loop over all candidates and append new permutation to all solutions `permute.go`
* Solution
    1. Accumulate through return statement append to list of solutions over candidates `solutionsWithN[j] = append(w, n)` `permute.go`

Combinations
* Generate = loop over candidates
    1. Loop over same constant list with `combination_sum.scala`
    2. Loop over list as arg with `combination_sum.go`, `linear_acc/letter_combinations.go`
    3. Consider one element at a time per recursive call `linear/letter_combinations.go`
* Separate function
    1. Yes `linear_acc/combination_sum.go`, `
* Order doesn't matter = how remove duplicates? 
    1. Set - problem is expensive `O(n)` and golang doesn't support slice keys.
    2. Only add if larger number e.g. `combination_sum.scala`
    3. Reduce the list of candidates with `[i:] combination_sum.go` or `digits[:len(digits)-1] linear/letter_combinations.go` 
* Base cases
    1. Empty list as base case `combination_sum.scala`
    2. Combination as base case, empty list if overshoot `combination_sum.go`
    3. Empty list as base case, but start combination if `len(1) linear/letter_combinations.go`
    4. Empty solution as base case for acc `"" linear_acc/combination_sum.go`
    5. Empty candidates and return `accString linear_acc/combination_sum.go`
* Solution
    1. Accumulate through return statement, append to list over candidates `combination_sum.scala`
    2. Accumulate through return statement, flat map over candidates, `combination_sum.go`
    3. Accumulate through return statement, append to list of solutions over candidates `linear/letter_combinations.go`
    4. Accumulate possible solution argument and final solutions to pointer `accString linear_acc/combination_sum.go` 
    5. Accumulate possible solution through argument and final solutions to return statement `accString linear_acc/combination_sum.go`


Subsets
* 