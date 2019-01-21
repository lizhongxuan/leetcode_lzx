package queue

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


func (q *Queue)Push(value interface{}) {
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

	q.top.preNode = node
	node.nextNode = q.top
	q.top = node

	return
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

func (s *Queue)PeekLast()interface{}{
	if s.botton == nil {
		return nil
	}
	return s.botton.value
}

func (s *Queue)PeekFirst()interface{}{
	if s.top == nil {
		return nil
	}
	return s.top.value
}

func (s *Queue)len()int{
	if s.top ==nil {
		return 0
	}
	step :=1
	node := s.top
	for node.nextNode != nil  {
		node = node.nextNode
		step +=1
	}
	return step
}