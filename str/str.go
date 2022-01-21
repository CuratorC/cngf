// Package str 字符串辅助方法
package str

import (
	"github.com/gertd/go-pluralize"
	"github.com/iancoleman/strcase"
)

// Plural 单词变复数
func Plural(word string) string {
	return pluralize.NewClient().Plural(word)
}

// Singular 单词转单数
func Singular(word string) string {
	return pluralize.NewClient().Singular(word)
}

// Snake 转为下划线蛇形命名法
func Snake(s string) string {
	return strcase.ToSnake(s)
}

// HigherCamel 转为大驼峰
func HigherCamel(s string) string {
	return strcase.ToCamel(s)
}

// LowerCamel 转为小驼峰
func LowerCamel(s string) string {
	return strcase.ToLowerCamel(s)
}
