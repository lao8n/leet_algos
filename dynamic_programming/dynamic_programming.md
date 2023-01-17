Types of algorithm
* Divide and conquer = independent (not overlapping) subproblems combined for solution
* Greedy algorithms = independent (not overlapping) subproblems that do not all have to be solved
* Dynamic programming = interdependent (overlapping) subproblems combined for solution

Two approaches to Dynamic Programming
* Bottom-up tabulation = start with the base cases and iterate up - adv: iteration faster than recursion
* Top-down memoization = start with case you want and recurse down - adv: ordering of recursion does not matter, or if not all subproblems need to be solved

When to use Dynamic programming
* max/min or number of ways 
* earlier decisions affect future decisions (i.e. cannot use greedy because of counter-example)

Dynamic programming checklist
* State function/data structure to give answer to the problem for every given state
* Recurrence relation to transition between states
* Base cases, so recursion/iteration ends