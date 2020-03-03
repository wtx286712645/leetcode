package leetcode

type ListNode struct {
    Val int
    Next *ListNode
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

    if l1 == nil {
        return l2
    }

    if l2 == nil {
        return l1
    }

	ll1 := getListLen(l1)
	ll2 := getListLen(l2)

	var short, long *ListNode
	if ll1 > ll2 {
		long = l1
		short = l2
	} else {
		short = l1
		long = l2
	}

	var n bool
	var lp, sp *ListNode
	lp = long
	sp = short
	for lp != nil {

		//两数先加
		if sp != nil {
			lp.Val += sp.Val
			sp = sp.Next
		}

		//是否有进位
		if n {
			lp.Val++
		}

		//如果有进位则设置进位标记，并减掉10
		if lp.Val >= 10 {
			lp.Val -= 10
			n = true
		} else {
			n = false
		}

		if lp.Next == nil && n {
			lp.Next = &ListNode{
				Val:1,
				Next:nil,
			}
			break
		}

		lp = lp.Next
	}

	return long
}

func getListLen(l *ListNode) int {

    ll := 0
	for l != nil {
		ll++
		l = l.Next
    }

    return ll
}
