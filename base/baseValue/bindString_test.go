package baseValue

import (
	"fmt"
	"testing"
)

func TestStrBindName(t *testing.T) {
	str := "{{name}}nuh{{wareNumber}}"
	args := map[string]any{"name": "hahah", "wareNumber": []float64{234.4, 23, 444}}
	spacing := []byte(" ")

	gotTemPs, err := StrBindName(str, args, spacing)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(gotTemPs)
}
