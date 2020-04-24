package main

import (
    "fmt"
)

type ListNode struct {
    Val int
    Next *ListNode
}

func main () {
    l1 := generateList([]int{2,4,3,2,4,3,2,4,3,2,4,3,2,4,3,2,4,3,2,4,3,2,4,3,2,4,3,2,4,3,2,4,3,2,4,3,2,4,3,2,4,3,2,4,3,2,4,3,2,4,3,2,4,3,2,4,3,2,4,3,9,1})
    l2 := generateList([]int{5,6,4,2,4,3,2,4,3,2,4,3,2,4,3,2,4,3,2,4,3,2,4,3,2,4,3,2,4,3,2,4,3,2,4,3,2,4,3,2,4,3,2,4,3,2,4,3,2,4,3,2,4,3,2,4,3,9,9,9,9})
    l3 := addTwoNumbers(l1, l2)
    fmt.Println(getNumbers(l3))
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    i, j := getNumbers(l1), getNumbers(l2)
    ret := plus(i, j)
    return generateList(reverse(ret))
}

func plus (v1, v2 []int) []int {
    v1, v2 = reverse(v1), reverse(v2)
    max, min := 0, 0
    var big []int
    if len(v1) > len(v2) {
        max, min, big = len(v1), len(v2), v1
    } else {
        max, min, big = len(v2), len(v1), v2
    }
    v3 := []int{}
    bit := false
    for i := 0; i < min; i++ {
        r := v1[i] + v2[i]
        if bit {
            r ++
            bit = false
        }
        if r >= 10 {
            bit = true
            r = r - 10
        }
        v3 = append(v3, r)
    }
    for j := min; j < max; j++ {
        if bit {
            bit = false
            r := big[j] + 1
            if r >= 10 {
                bit = true
                r = 0
            }
            v3 = append(v3, r)
        } else {
            v3 = append(v3, big[j])
        }
    }
    if bit {
        v3 = append(v3, 1)
    }
    return v3
}

func reverse (v []int) []int{
    v2 := []int{}
    for i := len(v) - 1; i >= 0; i -- {
        v2 = append(v2, v[i])
    }
    return v2
}

func getNumbers(l *ListNode) []int {
    i := []int{}
    for l != nil {
        i = append(i, l.Val)
        l = l.Next
    }
    return i
}

func generateList(val []int) *ListNode {
    l := &ListNode{
        Val: val[0],
    }
    m := l
    for i := 1; i < len(val); i++ {
        m.Next = &ListNode{
            Val: val[i],
        }
        m = m.Next
    }
    return l
}
