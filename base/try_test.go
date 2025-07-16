package base

import (
	"fmt"
	"testing"
)

func TestGuessProjectPathPrefix(t *testing.T) {
	got := guessProjectPathPrefix()
	fmt.Println(got)
}

func TestTry(t *testing.T) {
	defer Try(func(err any) {
		fmt.Println("错误输出：", err)
	})
	Iuyt()
}

func Iuyt() {
	for i := 0; i < 10; i++ {
		fmt.Println(1 / i)
	}
}
