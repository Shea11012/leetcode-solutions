package easy

import (
	"fmt"
	"testing"
)

func Test_GroupAnagrams(t *testing.T) {
	//testcase := []string{"cab","tin","pew","duh","may","ill","buy","bar","max","doc"}
	testcase := []string{"eat","tea","tan","ate","nat","bat"}
	v := groupAnagrams(testcase)
	fmt.Printf("%v",v)
}
