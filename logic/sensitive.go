package logic

import (
	"chatroom/global"
	"strings"
)

// FilterSensitive 敏感词过滤器
func FilterSensitive(content string) string {
	for _, word := range global.SensitiveWords {
		content = strings.ReplaceAll(content, word, "**")
	}
	return content
}
