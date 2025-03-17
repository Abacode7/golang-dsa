package leetcode

 type ListNode struct{
	Val int
	Next *ListNode
 }


/**
QUESTION: 21. Merge Two Sorted Lists
TAGS: Easy
*/
 func MergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    mergedList := new(ListNode)
    mergedListRef := mergedList

    for list1 != nil || list2 != nil {
        if list1 != nil && list2 != nil {
            if list1.Val < list2.Val {
                mergedList.Next = list1
                list1 = list1.Next
            }else{
                mergedList.Next = list2
                list2 = list2.Next
            }
        }else if list1 != nil {
            mergedList.Next = list1
            list1 = list1.Next
        }else {
            mergedList.Next = list2
            list2 = list2.Next
        }
        mergedList = mergedList.Next
    }
    return mergedListRef.Next
}


/**
QUESTION: 141. Linked List Cycle
TAGS: Easy
*/
func HasCycle(head *ListNode) bool {
    set := make(map[*ListNode]bool)

    for head != nil {
        if _, ok := set[head]; ok {return true}
        set[head] = true
        head = head.Next
    }
    return false
}

func FloydHasCycle(head *ListNode) bool {
    slow, fast := head, head
    for fast != nil && fast.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
        if slow == fast {return true}
    }
    return false
}



/**
QUESTION: 160. Intersection of Two Linked Lists
TAGS: Easy
*/
func GetIntersectionNode(headA, headB *ListNode) *ListNode {
    set := make(map[*ListNode]bool)

    for node := headA; node != nil; node = node.Next {
        set[node] = true
    }

    for node := headB; node != nil; node = node.Next {
        if _, ok := set[node]; ok {
            return node
        }
    }
    return nil
}