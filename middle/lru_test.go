package middle

import (
	"reflect"
	"testing"
)

func Test_AddFirst(t *testing.T) {
	d := constructorDoublelist()
	d.addFirst(&node{key:1,value:2})
	d.addFirst(&node{key:2,value:3})
	d.addFirst(&node{key:3,value:4})
	if d.Size() != 3 {
		t.Errorf("错误的大小：%d，期望大小：%d",d.Size(),3)
	}
}

func Test_RemoveLast(t *testing.T) {
	d := constructorDoublelist()
	d.addFirst(&node{key:1,value:2})
	d.addFirst(&node{key:2,value:3})
	d.addFirst(&node{key:3,value:4})
	d.removeLast()
	d.removeLast()
	d.removeLast()
	d.removeLast()

	want := 0
	if d.rear.prev.value != want {
		t.Errorf("期望获得值为%d，实际获得值为%d",want,d.rear.prev.value)
	}
}

func Test_constructor(t *testing.T) {
	lruC := constructor(3)
	k := reflect.TypeOf(lruC).Elem().Name()
	if k != "lruCache" {
		t.Errorf("实例错误 %s,want %s",k,"lruCache")
	}
}

func Test_Put(t *testing.T) {
	lruc := constructor(3)
	lruc.put(1,1)
	lruc.put(2,2)
	lruc.put(3,3)
}

func Test_Get(t *testing.T) {
	lruc := constructor(3)
	lruc.put(1,1)
	lruc.put(2,2)
	lruc.put(3,3)
	lruc.put(4,4)

	lruc.get(2)
	lruc.put(5,5)
	/*want := 3
	if lruc.head.value != want {
		t.Errorf("实际head 上的值为:%d,期望的值为：%d",lruc.head.value,want)
	}*/
}

func Test_LRUCache(t *testing.T) {
	lruc := constructor(2)
	// [[2],[1,1],[2,2],[1],[3,3],[2],[4,4],[1],[3],[4]]
	lruc.put(1,1)
	lruc.put(2,2)
	lruc.get(1)
	lruc.put(3,3)
	lruc.get(2)
	lruc.put(4,4)
	lruc.get(1)
	lruc.get(3)
	lruc.get(4)
}