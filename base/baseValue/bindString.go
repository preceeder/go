package baseValue

import (
	"errors"
	"fmt"
	"strings"
)

func StrBindName(str string, args map[string]any, spacing []byte) (string, error) {
	var builder strings.Builder
	offset := 0

	for {
		start := strings.Index(str[offset:], "{{")
		if start == -1 {
			builder.WriteString(str[offset:])
			break
		}
		start += offset

		end := strings.Index(str[start:], "}}")
		if end == -1 {
			builder.WriteString(str[offset:])
			break
		}
		end += start

		// 写入前部分文本
		builder.WriteString(str[offset:start])

		// 提取变量名
		key := strings.TrimSpace(str[start+2 : end])

		valStr, err := AnyToString(args[key], spacing)
		if err != nil {
			// 未找到变量，保留原样
			valStr = "{{" + key + "}}"
			return valStr, errors.New(fmt.Sprintf("not fond %s", key))
		}
		builder.WriteString(valStr)

		offset = end + 2
	}

	return builder.String(), nil
}
