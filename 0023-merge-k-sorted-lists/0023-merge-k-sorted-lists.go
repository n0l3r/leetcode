// Source: https://leetcode.com/problems/merge-k-sorted-lists/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
	dummy := &ListNode{}
	cur := dummy
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
	for _, list := range lists {
		if list != nil {
			heap.Push(&pq, list)
		}
	}
	for pq.Len() > 0 {
		node := heap.Pop(&pq).(*ListNode)
		cur.Next = node
		cur = cur.Next
		if node.Next != nil {
			heap.Push(&pq, node.Next)
		}
	}
	return dummy.Next
}

type PriorityQueue []*ListNode

func (pq PriorityQueue) Init() {
	n := len(pq)
	for i := n/2 - 1; i >= 0; i-- {
		pq.down(i, n)
	}
}

func (pq PriorityQueue) down(i, n int) {
	for {
		l := 2*i + 1
		r := 2*i + 2
		j := i
		if l < n && pq[j].Val > pq[l].Val {
			j = l
		}
		if r < n && pq[j].Val > pq[r].Val {
			j = r
		}
		if j == i {
			break
		}
		pq.Swap(i, j)
		i = j
	}
}

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Val < pq[j].Val
}

func (pq PriorityQueue) Swap(i, j int) {

	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	node := x.(*ListNode)
	*pq = append(*pq, node)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	node := old[n-1]
	*pq = old[0 : n-1]
	return node
}

