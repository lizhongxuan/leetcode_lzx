package main

import (
	"fmt"
)

func main() {
	//q :=  queue.NewQueue()
	//q.PushLast(1)
	//fmt.Println(q.PeekFirst().(int))
	//fmt.Println(q.PeekFirst().(int))
	fmt.Println(maxSlidingWindow([]int{1,3,1,2,0,5},3))
}

//答案
func maxSlidingWindow2(nums []int, k int) []int {
	if nums == nil {
		return nil
	}
	var res, window []int
	for i,val:= range nums {
		//移除超出左边界的元素
		if i-k >= 0 && window[0] <= i-k {
			window = window[1:]
		}
		//移除右侧小于等于val的元素,这样保持最左侧为最大元素
		for  len(window) > 0 && nums[window[len(window)-1]] <= val {
			window = window[:len(window)-1]
		}

		window = append(window,i) //入窗
		if i >= k-1{
			res = append(res, nums[window[0]])
		}
	}
	return res
}

func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) < k {
		return nil
	}
	q := NewQueue()
	res := []int{}
	for i:=0;i<len(nums) ;i++  {
		fmt.Println("i:",i)
		if i>=k && i-k >=q.PeekFirst().(int) {
			q.PopFirst()

			//检查队列中的最大的值是否在最前面,若不是,则把之前的数值排除
		}


		for q.PeekLast() != nil && nums[i]>=nums[q.PeekLast().(int)]  {
			q.Pop()
		}

		q.PushLast(i)

		fmt.Println("i:",i,"  k-1:",k-1)
		if i>=k-1 {
			res = append(res,nums[q.PeekFirst().(int)])
		}
	}

	return res
}

type Queue struct {
	top *queueNode
	botton *queueNode
}

type queueNode struct {
	value interface{}
	nextNode *queueNode
	preNode *queueNode
}

func NewQueue()*Queue  {
	return &Queue{
		top:nil,
		botton:nil,
	}
}
func (q *Queue)PushLast(value interface{})  {
	if q.top == nil {
		node := &queueNode{
			value:value,
		}
		q.top = node
		q.botton = node
		return
	}
	node := &queueNode{
		value:value,
	}
	q.botton.nextNode = node
	node.preNode = q.botton
	q.botton = node
	return
}
func (s *Queue)PopFirst()interface{}  {

	if s.top == nil  {
		return nil
	}
	node := s.top

	if s.top.nextNode == nil {
		s.top = nil
		s.botton = nil
	}else {
		s.top = s.top.nextNode
		s.top.preNode = nil
	}

	return node.value

}

func (s *Queue)isEmpty()bool{
	if s.botton == nil {
		return true
	}
	return false
}
func (s *Queue)PeekFirst()interface{}{
	if s.top == nil {
		return nil
	}
	return s.top.value
}

func (s *Queue)PeekLast()interface{}{
	if s.botton == nil {
		return nil
	}
	return s.botton.value
}

func (s *Queue)Pop()interface{}  {

	if s.botton == nil  {
		return nil
	}
	node := s.botton

	if s.botton.preNode == nil {
		s.top = nil
		s.botton = nil
	}else {
		s.botton = s.botton.preNode
		s.botton.nextNode = nil
	}

	return node.value

}