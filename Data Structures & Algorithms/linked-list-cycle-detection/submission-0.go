func hasCycle(head *ListNode) bool {
    slowNode := head // current
	fastNode := head // current
	for fastNode != nil && fastNode.Next != nil { // look into the future 2 Nodes
		slowNode = slowNode.Next
		fastNode = fastNode.Next.Next
		if slowNode == fastNode {
			return true
		}
		
	}
	return false
}
