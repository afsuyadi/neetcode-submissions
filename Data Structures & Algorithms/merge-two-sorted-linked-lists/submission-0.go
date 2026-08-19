
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	sortedList1 := list1 //sortedLists(list1)
	sortedList2 := list2 //sortedLists(list2)
	dummy := &ListNode{}
	current := dummy
	for sortedList1 != nil && sortedList2 != nil {
		if sortedList1.Val < sortedList2.Val {
			current.Next = sortedList1
			sortedList1 = sortedList1.Next
		} else {
			current.Next = sortedList2
			sortedList2 = sortedList2.Next
		}	
		current = current.Next	
	}
	
	// 0st loop
	// list1 = [4,2,1]
	// list2 = [5,3,1]
	// current = []
	// 1st loop
	// list1 = [2,1]
	// list2 = [5,3,1]
	// current = [4]

	// 2nd loop
	// list1 = [1]
	// list2 = [5,3,1]
	// current = [4,2]

	// 3rd loop
	// list1 = []
	// list2 = [5,3,1]
	// current = [4,2,1]
	if sortedList1 != nil {
		current.Next = sortedList1
	} else {
		current.Next = sortedList2
	}
	// dummy.Next = current
	return dummy.Next
}
