/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode = nil
	curr := head
	// next := curr.Next
    for curr != nil {
		// next = prev
		next := curr.Next
		curr.Next = prev
		// prev += 1
		prev = curr
		// curr = curr.Next
		curr = next
	}
	head = prev
	return head
}
