package easy

import (
	"fmt"
	"testing"
)

func Test_BackSpaceCompare(t *testing.T) {
	testcase := []string{"xywrrmp",
		"xywrrmu#p"}
	fmt.Printf("%v",backspaceCompare(testcase[0],testcase[1]))
}
