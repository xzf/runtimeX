package runtimeX

import (
	"github.com/xzf/strX"
	"runtime/debug"
	"strings"
)

func CallerLast() string {
	return Caller(2)
}

func CallerThis() string {
	return Caller(1)
}

func Caller(skip int) string {
	stack := string(debug.Stack())
	for i := 0; i < (skip+3)*2; i++ {
		stack = strX.SubAfter(stack, "\n")
	}
	stack = strings.TrimLeft(stack, "\t")
	stack = strX.SubBefore(stack, " ")
	if strings.Contains(stack, "\n") {
		stack = strX.SubBefore(stack, "\n")
	}
	return stack
}

func Stack() string {
	return string(debug.Stack())
}
