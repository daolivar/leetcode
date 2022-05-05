/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

/*
   Algorithm:
       traverse list
       get the int val of the nodes for list 1 & list2; 1->2->3 = 321 4->5->6 = 654
       add them (654+321) = 975
       then break that number down from a int to list; 975 -> 5->7->9
       return new list

*/
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	tmp := listToInt(l1) + listToInt(l2)
	return intToList(tmp)
}

func listToInt(l1 *ListNode) int {
	sum := 0
	currentNode := l1
	i := 1
	for currentNode != nil {
		sum += (currentNode.Val * i)
		i *= 10
		currentNode = currentNode.Next
	}
	return sum
}

// FIXME: // implement func to create new list
func intToList(n int) *ListNode {
	var res *ListNod
	for n != 0 {
		remainder := n % 10
		// implement func to create node
		fmt.Println(remainder)
		n /= 10
	}
	return res
}
