/**
 * Definition for singly-linked list.
 * class ListNode(_x: Int = 0, _next: ListNode = null) {
 *   var next: ListNode = _next
 *   var x: Int = _x
 * }
 */
object Solution {
    def reverseList(head: ListNode): ListNode = {
        def recursiveReverseList(head : ListNode, acc : ListNode): ListNode = {
            head match {
                case null => acc
                case head => recursiveReverseList(head.next, ListNode(head.x, acc))
            }
        }
        recursiveReverseList(head, null)
    }
}