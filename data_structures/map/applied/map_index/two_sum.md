Question https://leetcode.com/problems/two-sum/

Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

You can return the answer in any order.

Solution
1. brute force 
two for loops but time complexity is O(n^2), space is O(1)
2. two passes of hash map of value:index
time complexity is O(n) although have two passes , space is O(n) for the 
items stored in the hash table
3. one-pass hash map of value: index
time complexity : O(n) with lookups for O(1), space is O(n), but only 
need one pass