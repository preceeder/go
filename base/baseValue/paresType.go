package baseValue

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

// 传入指定的 类型， 我对应的字符串    使用场景 已知一个数据的 类型， 和一个已知的 字符串  且可以传化为那类型， 返回对应类型的reflect.value
func ParseValueByType(typ reflect.Type, str string) (reflect.Value, error) {
	typ = indirectType(typ)
	str = strings.TrimSpace(str)

	switch typ.Kind() {
	case reflect.String:
		return reflect.ValueOf(unquote(str)).Convert(typ), nil

	case reflect.Bool:
		v, err := strconv.ParseBool(str)
		return reflect.ValueOf(v).Convert(typ), err

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		v, err := strconv.ParseInt(str, 10, 64)
		return reflect.ValueOf(v).Convert(typ), err

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		v, err := strconv.ParseUint(str, 10, 64)
		return reflect.ValueOf(v).Convert(typ), err

	case reflect.Float32, reflect.Float64:
		v, err := strconv.ParseFloat(str, 64)
		return reflect.ValueOf(v).Convert(typ), err

	case reflect.Slice:
		return parseSlice(str, typ)

	case reflect.Map:
		return parseMap(str, typ)

	case reflect.Ptr:
		elemVal, err := ParseValueByType(typ.Elem(), str)
		if err != nil {
			return reflect.Value{}, err
		}
		ptr := reflect.New(typ.Elem())
		ptr.Elem().Set(elemVal)
		return ptr, nil

	default:
		return reflect.Value{}, fmt.Errorf("不支持的类型: %s", typ.Kind())
	}
}

func indirectType(t reflect.Type) reflect.Type {
	for t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	return t
}

func unquote(s string) string {
	if len(s) >= 2 && s[0] == '"' && s[len(s)-1] == '"' {
		return s[1 : len(s)-1]
	}
	return s
}

func parseSlice(str string, typ reflect.Type) (reflect.Value, error) {
	if str == "null" || str == "" {
		return reflect.MakeSlice(typ, 0, 0), nil
	}
	if str[0] != '[' || str[len(str)-1] != ']' {
		return reflect.Value{}, fmt.Errorf("无效的 slice 字符串: %s", str)
	}
	elements := SplitJson(str[1 : len(str)-1]) // 你需实现此函数
	slice := reflect.MakeSlice(typ, len(elements), len(elements))
	for i, part := range elements {
		elemVal, err := ParseValueByType(typ.Elem(), part)
		if err != nil {
			return reflect.Value{}, err
		}
		slice.Index(i).Set(elemVal)
	}
	return slice, nil
}

func parseMap(str string, typ reflect.Type) (reflect.Value, error) {
	if str == "null" || str == "" {
		return reflect.MakeMapWithSize(typ, 0), nil
	}
	if str[0] != '{' || str[len(str)-1] != '}' {
		return reflect.Value{}, fmt.Errorf("无效的 map 字符串: %s", str)
	}
	pairs := SplitJson(str[1 : len(str)-1]) // 支持 key:value 的字符串切割
	m := reflect.MakeMapWithSize(typ, len(pairs))
	for _, kv := range pairs {
		parts := strings.SplitN(kv, ":", 2)
		if len(parts) != 2 {
			return reflect.Value{}, fmt.Errorf("无效键值对: %s", kv)
		}
		k, err := ParseValueByType(typ.Key(), parts[0])
		if err != nil {
			return reflect.Value{}, fmt.Errorf("map key 解析失败: %w", err)
		}
		v, err := ParseValueByType(typ.Elem(), parts[1])
		if err != nil {
			return reflect.Value{}, fmt.Errorf("map value 解析失败: %w", err)
		}
		m.SetMapIndex(k, v)
	}
	return m, nil
}

func SplitJson(input string) []string {
	var result []string
	var buf strings.Builder

	depth := 0
	inString := false
	escape := false

	for i := 0; i < len(input); i++ {
		ch := input[i]

		if inString {
			buf.WriteByte(ch)

			if escape {
				escape = false
			} else if ch == '\\' {
				escape = true
			} else if ch == '"' {
				inString = false
			}
			continue
		}

		switch ch {
		case '"':
			inString = true
			buf.WriteByte(ch)

		case '{', '[':
			depth++
			buf.WriteByte(ch)

		case '}', ']':
			depth--
			buf.WriteByte(ch)

		case ',':
			if depth == 0 {
				// 当前层级是顶层，认为是分隔符
				result = append(result, strings.TrimSpace(buf.String()))
				buf.Reset()
			} else {
				buf.WriteByte(ch)
			}

		default:
			buf.WriteByte(ch)
		}
	}

	if buf.Len() > 0 {
		result = append(result, strings.TrimSpace(buf.String()))
	}

	return result
}
