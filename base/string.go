package base

import (
	"math/rand"
	"strconv"
	"time"
	"unicode/utf8"
)

// 随机字符串
var letters = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
var lettersInt = []rune("0123456789")

func RandStr(str_len int) string {
	rand_bytes := make([]rune, str_len)
	for i := range rand_bytes {
		rand_bytes[i] = letters[rand.Intn(len(letters))]
	}
	return string(rand_bytes)
}

func RandStrInt(str_len int) string {
	rand_bytes := make([]rune, str_len)
	for i := range rand_bytes {
		rand_bytes[i] = letters[rand.Intn(len(lettersInt))]
	}
	return string(rand_bytes)
}

// 最少13位
func GenterWithoutRepetitionStr(strl int) string {
	tim := strconv.FormatInt(time.Now().UnixMilli(), 10)
	if strl < 13 {
		return tim
	}
	netstr := RandStr(strl - 13)
	return tim + netstr
}

// ReplaceXin 使用 * 替换字符串中间部分
func ReplaceXin(str string) string {
	lens := StringLen(str)

	index := int(lens / 3)
	end := index * 2
	var istr string
	for range index {
		istr += "*"
	}
	return str[:index] + istr + str[end:]
}

// 获取字符串的 长度
func StringLen(str string) int {
	return utf8.RuneCountInString(str)
}

// 获取字符串的 字串, 这里可以防止中文截到一半的情况
func GetSubString(str string, start, end int) string {
	sn := []rune(str)
	return string(sn[start:end])
}
