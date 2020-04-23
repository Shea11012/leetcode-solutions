package easy

type node struct {
	data int
	next *node
	min int
}

type minStack struct {
	head *node
	len int
}

func constructor() minStack {
	return minStack{}
}

func (this *minStack) push(x int) {
	n := &node{data:x}
	if this.head == nil {
		this.head = n
		this.head.min = x
		return
	}
	if x < this.head.min {
		n.min = x
	} else {
		n.min = this.head.min
	}

	n.next = this.head
	this.head = n
}

func (this *minStack) pop() {
	if this.head != nil {
		this.head = this.head.next
	}
}

func (this *minStack) top() int {
	if this.head != nil {
		return this.head.data
	}

	return -1
}

func (this *minStack) getMin() int {
	if this.head != nil {
		return this.head.min
	}

	return -1
}