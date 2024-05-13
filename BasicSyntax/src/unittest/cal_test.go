package unittest

import (
	"fmt"
	"testing"
)

func TestAddUer1(t *testing.T) { //必须以Test开头和程序必须以_test结尾
	res := addUpper1(10)
	if res != 55 {
		t.Fatalf("addUpper执行错误，期望值：%v，实际值:%v", 55, res)
	}
	t.Logf("addUpper执行正确。。")
}
func TestHello(t *testing.T) { //以Test开头都会被调用
	fmt.Println("testhello被调用")
}
