package middle

type nodef struct {
	prev *nodef
	data	int
	next *nodef
}

type listF struct {
	head *nodef
	tail *nodef
}

func newListFer() *listF {
	//node := &nodef{}
	head,tail := &nodef{},&nodef{}
	head.next = tail // 初始化没有任何元素时，就是head->tail ，tail -> head
	tail.prev = head
	listF := &listF{
		head: head,
		tail: tail,
	}

	return listF
}

func (l *listF) add(node *nodef) {
	node.prev = l.head
	node.next = l.head.next
	l.head.next.prev = node // 将 tail 的前置节点改为 node
	l.head.next = node // 将 l.head 的后继节点改为 node

	/*if l.head.next == nil {
		l.head.next = node
	}

	node.prev = l.tail
	l.tail.next = node
	l.tail = node*/
}

func (l *listF) remove(node *nodef) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (l *listF) isEmpty() bool {
	if l.head == l.tail.prev {
		return true
	}

	return false
}

type FirstUnique struct {
	hmap map[int]*nodef
	llist *listF
}

func Constructor(nums []int) FirstUnique {
	firstUnique := FirstUnique{
		hmap:  map[int]*nodef{},
		llist: newListFer(),
	}

	for i := range nums {
		firstUnique.Add(nums[i])
	}

	return firstUnique
}

func (this *FirstUnique) ShowFirstUnique() int {
	if !this.llist.isEmpty() {
		return this.llist.tail.prev.data
	}
	return -1
}

func (this *FirstUnique) Add(value int) {
	node := &nodef{data:value}
	v,ok := this.hmap[value]
	if ok {
		if v != nil {
			this.llist.remove(v)
		}
		this.hmap[value] = nil
	} else {
		this.llist.add(node)
		this.hmap[value] = node
	}
}