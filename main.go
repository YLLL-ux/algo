package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	A := headA
	B := headB
	for A != B { // A指针与B指针重合位置，即为返回值
		if A != nil {
			A = A.Next
		} else { // A遍历完后，A指针指向headB
			A = headB
		}

		if B != nil {
			B = B.Next
		} else { // B遍历完后，B指针指向headA
			B = headA
		}
	}

	return A
}
