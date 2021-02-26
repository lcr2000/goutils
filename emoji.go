package goutils

import (
	"strings"
	"unicode/utf8"
)

// EmojiFilter过滤掉传入的字符串content中的Emoji
func EmojiFilter(content string) (newContent string) {
	for _, value := range content {
		// 解码 string中 UTF-8 编码序列，返回该码值和长度（比如中文字节长度3、英文1）
		_, size := utf8.DecodeRuneInString(string(value))
		if size >= 4 {
			continue
		}
		newContent += string(value)
	}
	return strings.Replace(newContent, " ", "", -1)
}
