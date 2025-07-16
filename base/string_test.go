package base

import (
	"fmt"
	"testing"
)

func TestGetSubString(t *testing.T) {
	//sub := GetSubString("你好呀", 1, 2)
	str := "你上课i大白菜似的次四粗俗彻底"
	fmt.Println(GetSubString(str, 1, 2))
	fmt.Println(str[1:2])
}
