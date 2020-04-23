package easy

import (
	"fmt"
	"testing"
)

func Test_CountingElements(t *testing.T) {
	testcase := []int{1,3,2,3,5,0}
	fmt.Println(countElements(testcase))
}
