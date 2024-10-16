package linkedlist

type SNode struct {
	value string
	next  *SNode
}

func NewSNode(value string, next *SNode) *SNode {
	return &SNode{value: value, next: next}
}

type SinglyLinkedList struct {
	head *SNode
}

func NewSinglyLinkedList() *SinglyLinkedList {
	return &SinglyLinkedList{}
}

/** Todo: GetAt, InsertAt, RemoveAt, GetHead, getTail, removeHead, removeTail
 */

func (r *SinglyLinkedList) GetHead() *SNode {
	return r.head
}

func (r *SinglyLinkedList) GetAt(index int) *SNode {
	if r.head == nil {
		return nil
	}

	position := 0
	c := r.head

	for position < index && c != nil {
		position += 1
		c = c.next
	}

	return c
}

func (r *SinglyLinkedList) InsertAt(index int, nodeToInsert *SNode) {
	/** Todo: To insert to a position, we need to find the node before that position
	 */
	if index == 0 {
		nodeToInsert.next = r.head
		r.head = nodeToInsert
		return
	}

	node := r.GetAt(index - 1)
	if node != nil {
		temp := node.next
		node.next = nodeToInsert
		nodeToInsert.next = temp
	}
}

func (r *SinglyLinkedList) RemoveAt(index int) *SNode {
	if index == 0 {
		if r.head != nil {
			nodeToRemove := r.head
			r.head = nodeToRemove.next
			nodeToRemove.next = nil
			return nodeToRemove
		} else {
			return nil
		}
	}

	node := r.GetAt(index - 1)
	if node != nil {
		nodeToRemove := node.next
		if nodeToRemove != nil {
			node.next = nodeToRemove.next
			nodeToRemove.next = nil
			return nodeToRemove
		}
	}
	return nil
}
