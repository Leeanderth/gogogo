package main

import "testing"

func TestAdd(t *testing.T) {
	sum := Add(1, -2)
	if sum < 0 {
		t.Errorf("测试失败")
		return
	}
	t.Logf("测试成功")

}
