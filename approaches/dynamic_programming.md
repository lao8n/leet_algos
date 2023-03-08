Types of algorithm
* Divide and conquer = independent (not overlapping) subproblems combined for solution
* Greedy algorithms = independent (not overlapping) subproblems that do not all have to be solved
* Dynamic programming = interdependent (overlapping) subproblems combined for solution

Two approaches to Dynamic Programming
* Bottom-up tabulation = start with the base cases and iterate up - adv: iteration faster than recursion
* Top-down memoization = start with case you want and recurse down - adv: ordering of recursion does not matter, or if not all subproblems need to be solved. Memoization requires immutable key however so no slices

Steps to convert top-down to bottom-up
1. initialize memoized slice of possible states
2. set base cases
3. for loops from bases cases
4. copy top-down logic to corresponding bottom up

When to use Dynamic programming
* max/min or number of ways 
* earlier decisions affect future decisions (i.e. cannot use greedy because of counter-example)

Dynamic programming checklist
* State function/data structure to give answer to the problem for every given state (e.g. index/indices, constraints)
* Recurrence relation to transition between states
* Base cases, so recursion/iteration ends

Implementations
* maybe dynamic programming solutions have been stored in either the recursion memoization sections or in data structures constant memory

Complexity
* O(n x F) where n is the number of possible states and F is the complexity of the recurrence relation.