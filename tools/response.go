package tools

import (
	"strings"
)

/*
 * 处理http响应的工具
 */

// RmUndefinedOfResponse 去除响应字符串中的undefined
func RmUndefinedOfResponse(str string) string {
	newStr := strings.Replace(str, ":undefined", `:"undefined"`, -1)
	return newStr
}
