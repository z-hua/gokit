package strs

import (
	"bytes"
	"strings"
)

// CamelCase 驼峰命名法，如 myName
func CamelCase(s string) string {
	bs := make([]byte, 0, len(s))
	for i := 0; i < len(s); i++ {
		if i == 0 {
			bs = append(bs, bytes.ToLower([]byte{s[i]})...)
			continue
		}
		if s[i] == '_' {
			i++
			bs = append(bs, bytes.ToUpper([]byte{s[i]})...)
			continue
		}
		bs = append(bs, s[i])
	}
	return string(bs)
}

// PascalCase 帕斯卡命名法，如 MyName
func PascalCase(s string) string {
	bs := make([]byte, 0, len(s))
	for i := 0; i < len(s); i++ {
		if i == 0 {
			bs = append(bs, bytes.ToUpper([]byte{s[i]})...)
			continue
		}
		if s[i] == '_' {
			i++
			if i >= len(s) {
				continue
			}
			bs = append(bs, bytes.ToUpper([]byte{s[i]})...)
			continue
		}
		bs = append(bs, s[i])
	}
	return string(bs)
}

// SnakeCase 蛇形命名法，如 my_name
func SnakeCase(s string) string {
	bs := make([]byte, 0, len(s))
	for i := 0; i < len(s); i++ {
		if s[i] >= 'A' && s[i] <= 'Z' && i > 0 && s[i-1] != '_' {
			bs = append(bs, '_')
		}
		bs = append(bs, s[i])
	}
	return strings.ToLower(string(bs))
}
