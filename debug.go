package wtf

import (
	"log"
	"runtime"
	"strings"
)

// Stack 堆栈追踪
func Stack(err error, args ...interface{}) {
	buf := make([]byte, 1<<10)
	runtime.Stack(buf, false)
	args = append([]interface{}{err}, args...)
	args = append(args, string(buf))
	argsFormat := strings.Repeat("%#v\r\n", len(args)-1) + "%+v\r\n"
	log.Printf(argsFormat, args...)
}
