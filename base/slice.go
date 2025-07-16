package base

import (
	"errors"
	"math/rand"
	"sort"
	"strconv"
	"strings"
)

// 快速版本：[]int → string
func IntSliceToString[T int32 | int | int64](slice []T, sep string) string {
	if len(slice) == 0 {
		return ""
	}
	var b strings.Builder
	tmp := make([]byte, 0, 32)

	tmp = strconv.AppendInt(tmp[:0], int64(slice[0]), 10)
	b.Write(tmp)

	for _, v := range slice[1:] {
		b.WriteString(sep)
		tmp = strconv.AppendInt(tmp[:0], int64(v), 10)
		b.Write(tmp)
	}
	return b.String()
}

// 快速版本：[]float64 → string（固定精度）
func FloatSliceToString[T float32 | float64](slice []T, sep string, prec int) string {
	if len(slice) == 0 {
		return ""
	}
	var b strings.Builder
	tmp := make([]byte, 0, 64)

	tmp = strconv.AppendFloat(tmp[:0], float64(slice[0]), 'f', prec, 64)
	b.Write(tmp)

	for _, v := range slice[1:] {
		b.WriteString(sep)
		tmp = strconv.AppendFloat(tmp[:0], float64(v), 'f', prec, 64)
		b.Write(tmp)
	}
	return b.String()
}

// 快速版本：[]string → string
func StringSliceToString(slice []string, sep string) string {
	return strings.Join(slice, sep)
}

// 通用泛型版本，支持任意类型 + 自定义格式化函数
func SliceToString[T any](slice []T, sep string, formatter func(T) string) string {
	if len(slice) == 0 {
		return ""
	}
	var b strings.Builder
	b.Grow(len(slice) * 8) // 预估容量

	b.WriteString(formatter(slice[0]))
	for _, v := range slice[1:] {
		b.WriteString(sep)
		b.WriteString(formatter(v))
	}
	return b.String()
}

// 对象数组 转化为 map对象
func StructSliceToStructMap[T ~[]E, E any, K comparable](datas T, f func(data E) (K, E)) map[K]E {
	var temp = map[K]E{}
	for _, v := range datas {
		k, value := f(v)
		temp[k] = value
	}
	return temp
}

// StructSliceToSliceValue 获取对象数组中某个字段的值的数组
func StructSliceToSliceValue[T ~[]E, E any, V any](datas T, f func(data E) V) []V {
	var temp = []V{}
	for _, v := range datas {
		temp = append(temp, f(v))
	}
	return temp
}

// SliceToMap
// two slice  to map
func SliceToMap[K comparable, V comparable](k []K, v []V) (map[K]V, error) {
	if len(k) > len(v) {
		return nil, errors.New("k len lt v len")
	}
	var res = make(map[K]V)
	for i, k := range k {
		res[k] = v[i]
	}
	return res, nil
}

// 特定类型的slice  转化为 any类型，   any 类型转化为 特定类型使用  cast.ToStringSlice()  就可以
func SliceConvertToAny[P comparable](src []P) []any {
	var dest = make([]any, len(src))
	for i, v := range src {
		dest[i] = v
	}
	return dest
}

// SlicesIntersect 求两个切片的交集
func SlicesIntersect[T comparable](a []T, b []T) []T {
	var inter []T
	mp := make(map[T]bool)

	for _, s := range a {
		if _, ok := mp[s]; !ok {
			mp[s] = true
		}
	}
	for _, s := range b {
		if _, ok := mp[s]; ok {
			inter = append(inter, s)
		}
	}

	return inter
}

// SlicesDiff 求两个切片的差集 a - b
func SlicesDiff[T comparable](a []T, b []T) []T {
	var diffArray []T
	temp := map[T]struct{}{}

	for _, val := range b {
		if _, ok := temp[val]; !ok {
			temp[val] = struct{}{}
		}
	}

	for _, val := range a {
		if _, ok := temp[val]; !ok {
			diffArray = append(diffArray, val)
		}
	}

	return diffArray
}

// SlicesUnique 切片去重
func SlicesUnique[T comparable](arr []T) []T {
	result := make([]T, 0, len(arr))
	temp := map[T]struct{}{}
	for i := 0; i < len(arr); i++ {
		if _, ok := temp[arr[i]]; ok != true {
			temp[arr[i]] = struct{}{}
			result = append(result, arr[i])
		}
	}
	return result
}

// 两个数组   一个数组按照另一个数组的顺序排序
// target 数组的顺序受 order 数组的顺序影响
// fo 方法返回 order 中关于 target 比较的数据
// f 方法 返回 trarget 中每个对象的关于order中数据的可比对象
// O 就是target, order 两个数组相关连的数据
func SlicesSortByAnotherArray[T comparable, V comparable, O comparable](target []T, order []V, fo func(V) O, f func(T) O) []T {
	// 构建顺序映射表，记录每个元素在参考数组中的顺序
	orderMap := make(map[O]int)
	for i, val := range order {
		orderMap[fo(val)] = i + 1 // 加 1 是为了区分未定义的值
	}

	// 自定义排序规则
	sort.Slice(target, func(i, j int) bool {
		posI, okI := orderMap[f(target[i])]
		posJ, okJ := orderMap[f(target[j])]

		if okI && okJ { // 如果两个元素都在顺序数组中，按顺序排序
			return posI < posJ
		} else if okI { // 仅一个元素在顺序数组中，在顺序数组中的优先
			return true
		} else if okJ {
			return false
		} else { // 两个元素都不在顺序数组中，按自然顺序排序, 之前是什么顺序，之后就按照这个顺序
			return false
		}
	})

	return target
}

// condition 返回true的就保留
func SlicesFilter[T any](input []T, condition func(T) bool) []T {
	var result []T
	for _, v := range input {
		if condition(v) {
			result = append(result, v)
		}
	}
	return result
}

// 从数组中随机获取几个数据， 不重复的数据
func SliceShuffleGetEleNumber[T any](input []T, num int) []T {
	keyMap := map[int]any{}
	i := 0
	length := len(input)
	if length < num {
		num = length
	}
	var rs = make([]T, num)
	for i < num {
		randomInt := rand.Intn(length)
		if _, ok := keyMap[randomInt]; ok {
			continue
		}
		keyMap[randomInt] = 0
		rs[i] = input[randomInt]
		i += 1
	}
	return rs
}

// c从尾部弹出数据
func PopBack[T any](s []T) ([]T, T, bool) {
	if len(s) == 0 {
		var zero T
		return s, zero, false
	}
	elem := s[len(s)-1]
	return s[:len(s)-1], elem, true
}

// 从头部弹出数据
func PopFront[T any](s []T) ([]T, T, bool) {
	if len(s) == 0 {
		var zero T
		return s, zero, false
	}
	elem := s[0]
	return s[1:], elem, true
}
