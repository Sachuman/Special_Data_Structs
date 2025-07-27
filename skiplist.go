package main

import (
	"math"
	"math/rand"
)

type Skiplist struct {
	Head *Node
}

type Node struct {
	Val   int
	Count int
	Next  []*Node
}

func Constructor() Skiplist {

	skipy := Skiplist{
		Head: &Node{
			Val:  math.MinInt,
			Next: []*Node{nil},
		},
	}
	return skipy
}

func (s *Skiplist) Search(target int) bool {
	cur := s.Head
	level := len(cur.Next) - 1
	for {
		if cur.Next[level] == nil || cur.Next[level].Val > target {
			if level == 0 {
				return false
			}
			level--
		} else if cur.Next[level].Val < target {
			cur = cur.Next[level]
			level = len(cur.Next) - 1

		} else {
			return true
		}

	}

}

func (s *Skiplist) Add(num int) {
	cur := s.Head
	pivots := []*Node{cur}
	level := len(cur.Next) - 1

	for {
		if cur.Next[level] == nil || cur.Next[level].Val > num {
			if level == 0 {
				break
			}
			pivots = append(pivots, cur)
			level--
		} else if cur.Next[level].Val < num {
			cur = cur.Next[level]
			level = len(cur.Next) - 1
		} else if cur.Next[level].Val == num {
			cur.Next[level].Count++
			return
		}
	}
	n := &Node{Val: num, Count: 1, Next: []*Node{cur.Next[0]}}
	cur.Next[0] = n

	currLevel := 1
	for rand.Intn(10)%2 == 1 && len(pivots) > 0 {
		pivot := pivots[len(pivots)-1]
		pivots = pivots[:len(pivots)-1]

		if currLevel < len(pivot.Next) {
			next := pivot.Next[currLevel]
			pivot.Next[currLevel] = n
			n.Next = append(n.Next, next)
		} else {
			pivot.Next = append(pivot.Next, n)
			n.Next = append(n.Next, nil)
		}
		currLevel++
	}

}

func (s *Skiplist) Erase(num int) bool {

	cur := s.Head
	level := len(cur.Next) - 1

	for {
		if cur.Next[level] == nil || cur.Next[level].Val > num {
			if level == 0 {
				return false
			}
			level--
		} else if cur.Next[level].Val < num {
			cur = cur.Next[level]
			level = len(cur.Next) - 1
		} else if cur.Next[level].Val == num {
			if cur.Next[level].Count > 1 {
				cur.Next[level].Count--
				return true
			}
			cur.Next[level] = cur.Next[level].Next[level]
			level--
			if level < 0 {
				return true
			}
		}
	}
}

/**
 * Your Skiplist object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Search(target);
 * obj.Add(num);
 * param_3 := obj.Erase(num);
 */
