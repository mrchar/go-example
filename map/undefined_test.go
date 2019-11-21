package main

import (
	"fmt"
	"testing"
)

// 测试获取字典中未定义的值
func TestAssign(t *testing.T) {
	Map := make(map[string]string)
	fmt.Printf("value: %#v\n", Map["Hello"])
}

func TestDelete(t *testing.T) {
	m := make(map[string]string)
	delete(m, "hello")
}
