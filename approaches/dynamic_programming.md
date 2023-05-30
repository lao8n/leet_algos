Types of algorithm
* Divide and conquer = independent (not overlapping) subproblems combined for solution
* Greedy algorithms = independent (not overlapping) subproblems that do not all have to be solved
* Dynamic programming = interdependent (overlapping) subproblems combined for solution

When to use dynamic programming
* Max/min or number of ways 
* Earlier decisions affect future decisions (i.e. cannot use greedy because of counter-example)
* Pathing problems with a constraint such as not visiting the same square twice to prevent moving backwards (o/w use BFS)

Two approaches to dynamic programming
* Bottom-up tabulation = start with the base cases and iterate up - adv: iteration faster than recursion
* Top-down memoization = start with case you want and recurse down - adv: ordering of recursion does not matter, or if not all subproblems need to be solved. Memoization requires immutable key however so no slices

Dynamic programming checklist
* State function/data structure to give answer to the problem for every given state (e.g. index/indices, constraints)
* Recurrence relation to transition between states
* Base cases, so recursion/iteration ends

Code examples
* The dynamic programming solutions have been stored in either the recursion memoization sections or in data structures constant memory

Complexity
* O(n x F) where n is the number of possible states and F is the complexity of the recurrence relation.

Tricks
* State reduction - can we reduce it further? e.g. house robber don't need boolean of whether robbed previous house just need clever recurrence relation
* Space reduction - only last few elements are needed

Kadane's algorithm = can find the maximum sum subarray (only useful if can have negative numbers) - single transaction
```
1. best = neg infinity
2. current = 0
3. for num in nums:
    * current = max(current + num, num) // if current < 0 then reset
    * best = max(best, current)
4. return best
```

Kadane examples
* Minimum price and Profit `max_profit.go`
* Max sub array `max_sub_array.go`, `max_profit.go`, `max_circ_subarray_sum.go`