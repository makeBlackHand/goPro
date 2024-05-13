package unittest

import (
	"testing"
)

func TestGetSub(t *testing.T) {
	res := sub(10, 5)
	if res != 5 {
		t.Fatalf("sub输出错误,期望值:%v，实际值:%v", 7, res)
	}
	t.Logf("sub输出正确。。")
}
