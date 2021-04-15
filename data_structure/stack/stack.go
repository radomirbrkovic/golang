// This is example of implementing stack data structure
package main


type Stack struct {
	Items []int
}

func (S *Stack) Push(v int) {
	S.Items = append([]int{v}, S.Items...)
}

func (S *Stack) Pop()  {
	if len(S.Items) > 1 {
		S.Items = S.Items[1:]
	} else {
		S.Items = make([]int, 0)
	}

}