package rdb

import (
	"fmt"
	"testing"
)

func Test_highPerfReplace(t *testing.T) {
	// 模板字符串，使用 {{name}} 格式
	template := []byte("Hello, {{name}}! You are {{age}} years old. Price: {{price}}. Active: {{active}}.")

	// 替换数据，类型为 map[string]any
	replacements := map[string]any{
		"name":   "Alice",
		"age":    30.398034,
		"price":  19.990,
		"active": 34.55,
	}

	// 调用模板替换函数
	result := highPerfReplace(template, replacements)

	// 输出替换结果
	fmt.Println(string(result))
}
