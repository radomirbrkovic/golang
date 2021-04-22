// This is example of implementing queue data structure
package main

type Queue struct {
	Items []int
}

func (Q *Queue) Push(v int) {
	Q.Items = append([]int{v}, Q.Items...)
}

func (Q *Queue) Pop()  {
	if len(Q.Items) > 1 {
		Q.Items = Q.Items[:len(Q.Items) - 1]
	} else {
		Q.Items = make([]int, 0)
	}

}