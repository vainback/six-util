package six

import (
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

type XStr string

func Str(s string) XStr {
	return XStr(s)
}

func Strings(ss ...string) XStr {
	return XStr(strings.Join(ss, ""))
}

func (s XStr) Len() int {
	return len(s)
}

func (s XStr) Size() int {
	return utf8.RuneCountInString(s.String())
}

func (s XStr) IsEmpty() bool {
	return s.Len() == 0
}

func (s XStr) Int() int {
	i, _ := strconv.Atoi(string(s))
	return i
}

func (s XStr) Int64() int64 {
	return int64(s.Int())
}

func (s XStr) Uint() uint {
	return uint(s.Int())
}

func (s XStr) Uint64() uint64 {
	return uint64(s.Int())
}

func (s XStr) Float() float64 {
	f, _ := strconv.ParseFloat(string(s), 64)
	return f
}

func (s XStr) Bool() bool {
	b, _ := strconv.ParseBool(string(s))
	return b
}

func (s XStr) String() string {
	return string(s)
}

func (s XStr) Bytes() []byte {
	return []byte(s)
}

// FirstUpper 首字母大写
func (s XStr) FirstUpper() XStr {
	if s.Len() == 0 {
		return s
	}
	runes := []rune(s)
	if unicode.IsLetter(runes[0]) {
		runes[0] = unicode.ToUpper(runes[0])
	}
	return XStr(runes)
}

// FirstLower 首字母小写
func (s XStr) FirstLower() XStr {
	if s.Len() == 0 {
		return s
	}
	runes := []rune(s)
	if unicode.IsLetter(runes[0]) {
		runes[0] = unicode.ToLower(runes[0])
	}
	return XStr(runes)
}

// Append 追加多个字符串到s尾部
func (s XStr) Append(ss ...string) XStr {
	if len(ss) == 0 {
		return s
	}

	var str strings.Builder
	str.WriteString(s.String())
	for i := 0; i < len(ss); i++ {
		str.WriteString(ss[i])
	}

	return XStr(str.String())
}

// Added 将 str 插入到 s 头部
func (s XStr) Added(str string) XStr {
	return XStr(strings.Join([]string{str, s.String()}, ""))
}

// Split 将s以sep字符分割为数组/切片
func (s XStr) Split(sep string) []string {
	return strings.Split(string(s), sep)
}

type TrimType int

const (
	TrimLeft TrimType = iota
	TrimRight
)

func (s XStr) Trim(cutest string, direction ...TrimType) string {
	if len(direction) == 1 {
		switch direction[0] {
		case TrimLeft:
			return strings.TrimLeft(string(s), cutest)
		case TrimRight:
			return strings.TrimRight(string(s), cutest)
		}
	}
	return strings.Trim(string(s), cutest)
}

func (s XStr) TrimSpace() string {
	return strings.TrimSpace(string(s))
}
