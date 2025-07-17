package common

import (
	"math"
	"strings"
)

// BaseTo26 10 进制转化为26 进制
func BaseTo26(data int) string {
	ld := []string{}
	for {
		number := data % 26
		ld = append(ld, string(rune(number+65)))
		data = int(math.Ceil(float64(data/26))) - 1
		if data < 0 {
			break
		} else if data <= 26 && data > 0 {
			ld = append(ld, string(rune(data+65)))
			break
		}
	}
	return strings.Join(ld, "")
}
