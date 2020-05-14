package middle

// 双向链表
type node struct {
	prev *node
	key int
	value int
	next *node
}

// doublelist 双向链表
type doublelist struct {
	head *node
	rear *node
	size int
}

func constructorDoublelist() *doublelist {
	head := &node{}
	rear := &node{}
	d := &doublelist{
		head:head,
		rear:rear,
	}
	d.head.next = rear
	d.rear.prev = head
	return d
}

// 在头部添加节点
func (d *doublelist) addFirst(v *node) {
	v.prev = d.head
	v.next = d.head.next
	d.head.next.prev = v
	d.head.next = v
	d.size++
}

// 删除一个节点
func (d *doublelist) remove(v *node) {
	v.prev.next = v.next
	v.next.prev = v.prev
	d.size--
}

// 在尾部删除节点
func (d *doublelist) removeLast() *node {
	if d.head == d.rear.prev {
		return nil
	}
	last := d.rear.prev
	d.remove(d.rear.prev)
	return last
}

func (d doublelist) Size() int {
	return d.size
}

type lruCache struct {
	vMap map[int]*node
	capacity int
	list *doublelist
}

func constructor(c int) *lruCache {
	return &lruCache{
		vMap: make(map[int]*node),
		capacity: c,
		list:constructorDoublelist(),
	}
}

func (l *lruCache) get (key int) int {
	v,ok := l.vMap[key]
	if !ok {
		return -1
	}

	l.put(key,v.value)
	return v.value
}

func (l *lruCache) put (key int,v int) {

	node := &node{
		prev:  nil,
		key:   key,
		value: v,
		next:  nil,
	}

	if n, ok := l.vMap[key]; ok {
		l.list.remove(n)
		delete(l.vMap,key)
	} else if l.capacity == l.list.Size() {
		n := l.list.removeLast()
		if n != nil {
			delete(l.vMap,n.key)
		}
	}

	l.list.addFirst(node)
	l.vMap[key] = node
}