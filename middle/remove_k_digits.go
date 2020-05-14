package middle

import (
	"container/list"
	"strings"
)

func removeKdigits(num string, k int) string {
	numLen := len(num)
	if numLen == k {
		return "0"
	}
	stack := list.New()
	for _,v := range num {
		if stack.Len() > 0 {
			e := stack.Back()
			value := e.Value.(int32)
			// stack 是否为空必须判断，第一个元素可能会被删除
			// k 限制删除元素
			// value 与 v 比较判断大小
			for stack.Len() > 0 && k > 0 && value > v {
				if e != nil {
					stack.Remove(e)
				}
				e = stack.Back()
				if e != nil {
					value = e.Value.(int32)
				}
				k--
			}
		}
		stack.PushBack(v)
	}

	// 有可能是一组 12345 这类的字符串，此时的 k 就没有被减少，所以需要从尾部删除
	for stack.Len() > 0 && k > 0 {
		e := stack.Back()
		if e != nil {
			stack.Remove(e)
		}
		k--
	}
	var strs strings.Builder
	// 移除前导零
	for e := stack.Front();e != nil && e.Value.(int32) == '0' && stack.Len() > 1; e = stack.Front(){
		stack.Remove(e)
	}

	for v := stack.Front();v != nil;v = v.Next() {
		strs.WriteString(string(v.Value.(int32)))
	}

	return strs.String()
}
