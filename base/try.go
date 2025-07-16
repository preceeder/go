package base

import (
	"fmt"
	"path/filepath"
	"runtime"
	"strings"
)

var ProductName string

func init() {
	ProductName = guessProjectPathPrefix()
}

// 自动推测调用代码中的项目目录关键名
func guessProjectPathPrefix() string {
	_, file, _, ok := runtime.Caller(1)
	if !ok {
		return ""
	}

	// 转换成统一的路径格式
	path := filepath.ToSlash(file)

	// 举例：/Users/you/go/myapp/pkg/utils/log.go
	// 截出 "myapp" 目录作为关键路径
	parts := strings.Split(path, "/")
	for i := len(parts) - 1; i >= 0; i-- {
		if parts[i] == "go" || parts[i] == "src" || parts[i] == "Users" || parts[i] == "home" {
			break
		}
		// 判断是否为实际项目目录名
		if len(parts[i]) > 2 && !strings.HasSuffix(parts[i], ".go") {
			return parts[i] // 比如返回 "myapp"
		}
	}
	return ""
}

func printMyStackAuto() string {
	pcs := make([]uintptr, 16)
	n := runtime.Callers(3, pcs)
	frames := runtime.CallersFrames(pcs[:n])

	for {
		frame, more := frames.Next()
		if strings.Contains(frame.File, ProductName) {
			return fmt.Sprintf("%s:%d (%s)", frame.File, frame.Line, frame.Function)
		}
		if !more {
			break
		}
	}
	return "unknown"
}

func Try(f func(err any)) {
	if err := recover(); err != nil {
		f(printMyStackAuto())
	}
}
