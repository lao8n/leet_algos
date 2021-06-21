/**
 * Definition for singly-linked list.
 * class ListNode(_x: Int = 0, _next: ListNode = null) {
 *   var next: ListNode = _next
 *   var x: Int = _x
 * }
 */

// brute force approach
// go through each linked list to build the number - multiplying by 10 for each step
// then add the two numbers and then convert back into a linked list
// problem is O(3n)

// single pass
// loop through both lists in parallel - carrying over
// problem because of ordering we need to accumulate a listnode as well
// building in the wrong order - we could reverse but that would cost O(n)
// so can we build it the opposite way - maybe we have to accumulate a number and then output afterwards

object Solution {
    def addTwoNumbers(l1: ListNode, l2: ListNode): ListNode = {
        def recursiveAddTwoNumbers(l1: ListNode, l2: ListNode, carryOver : Int): ListNode = {
            (l1, l2) match {
                // base case
                case (null, null) if carryOver > 0 => ListNode(carryOver, null)
                case (null, null) => null
                // recursive case
                case (l1, null) => 
                    ListNode((l1.x + carryOver) % 10, 
                    recursiveAddTwoNumbers(l1.next, null,(l1.x + carryOver) / 10))
                case (null, l2) => 
                    ListNode((l2.x + carryOver) % 10, 
                    recursiveAddTwoNumbers(null, l2.next,(l2.x + carryOver) / 10))
                case (l1, l2) => 
                    ListNode((l1.x + l2.x + carryOver) % 10, 
                    recursiveAddTwoNumbers(l1.next, l2.next,(l1.x + l2.x + carryOver) / 10))
            }
        }
        recursiveAddTwoNumbers(l1, l2, 0)
    }
}