package baseValue

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

type Method struct {
	Servers map[string]reflect.Method
	Rcvr    reflect.Value
	Typ     reflect.Type
}

// MakeService rep 的 传入 指针 非指针都可以
// 解析对象的 方法 可以通过 map的形式执行
func MakeService(rep interface{}) *Method {
	ser := Method{}
	ser.Typ = reflect.TypeOf(rep)
	ser.Rcvr = reflect.ValueOf(rep)
	//name := reflect.Indirect(ser.Rcvr).Type().Name()
	ser.Servers = map[string]reflect.Method{}
	for i := 0; i < ser.Typ.NumMethod(); i++ {
		method := ser.Typ.Method(i)
		mname := method.Name // string
		ser.Servers[mname] = method
	}

	return &ser
}

// 任意数据转化为 字符串
func AnyToString(value any, spacing []byte) (string, error) {
	if value == nil {
		return "", errors.New("参数为 nil")
	}

	val := reflect.ValueOf(value)
	val = reflect.Indirect(val) // 解除指针
	typ := val.Type()

	sep := string(spacing)

	switch val.Kind() {
	case reflect.String:
		return val.String(), nil

	case reflect.Bool:
		return fmt.Sprintf("%t", val.Bool()), nil

	case reflect.Float32, reflect.Float64:
		// 这里不使用 fmt.Sprintf("%f", val.Float())  是因为需要保持原本的数据格式， 不然会在后面多处0的
		return fmt.Sprintf("%v", val.Interface()), nil

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return fmt.Sprintf("%d", val.Int()), nil

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return fmt.Sprintf("%d", val.Uint()), nil

	case reflect.Slice, reflect.Array:
		if val.Len() == 0 {
			return "", nil
		}
		var parts []string
		for i := 0; i < val.Len(); i++ {
			elemStr, err := AnyToString(val.Index(i).Interface(), spacing)
			if err != nil {
				return "", err
			}
			parts = append(parts, elemStr)
		}
		return strings.Join(parts, sep), nil

	case reflect.Map:
		if val.Len() == 0 || val.IsNil() {
			return "", nil
		}
		var parts []string
		for _, key := range val.MapKeys() {
			kStr, err := AnyToString(key.Interface(), spacing)
			if err != nil {
				return "", err
			}
			vStr, err := AnyToString(val.MapIndex(key).Interface(), spacing)
			if err != nil {
				return "", err
			}
			parts = append(parts, fmt.Sprintf("%s=%s", kStr, vStr))
		}
		return strings.Join(parts, sep), nil

	case reflect.Struct:
		var parts []string
		for i := 0; i < val.NumField(); i++ {
			field := val.Field(i)
			fieldType := typ.Field(i)

			// 只处理导出字段
			if !field.CanInterface() {
				continue
			}

			key := fieldType.Name
			if jsonTag := fieldType.Tag.Get("json"); jsonTag != "" {
				tagParts := strings.Split(jsonTag, ",")
				if tagParts[0] != "-" && tagParts[0] != "" {
					key = tagParts[0]
				}
			}

			valStr, err := AnyToString(field.Interface(), spacing)
			if err != nil {
				return "", err
			}
			parts = append(parts, fmt.Sprintf("\"%s\":%s", key, valStr))
		}
		return strings.Join(parts, sep), nil

	case reflect.Interface:
		return AnyToString(val.Interface(), spacing)

	default:
		return fmt.Sprintf("%v", value), nil
	}
}

func RunFunc(object interface{}, methodName string, args ...interface{}) []reflect.Value {
	inputs := make([]reflect.Value, len(args))
	for i, _ := range args {
		inputs[i] = reflect.ValueOf(args[i])
	}
	//动态调用方法
	return reflect.ValueOf(object).MethodByName(methodName).Call(inputs)
}

func ReflectBaseType(t reflect.Type) reflect.Type {
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	return t
}

func IsTargetType(t reflect.Type, expected reflect.Kind) (reflect.Type, error) {
	t = ReflectBaseType(t)
	if t.Kind() != expected {
		return nil, fmt.Errorf("expected %s but got %s", expected, t.Kind())
	}
	return t, nil
}

type CanCalcType interface {
	~int | ~uint8 | ~int16 | ~int32 | ~int64 | ~float32 | ~float64
}

func GetCalcAttr[d CanCalcType](s interface{}, fieldName string) d {
	v := reflect.ValueOf(s)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	var zore d
	if v.Kind() != reflect.Struct {
		return zore
	}

	f := v.FieldByName(fieldName)
	if !f.IsValid() {
		return zore
	}

	return f.Interface().(d)
}
