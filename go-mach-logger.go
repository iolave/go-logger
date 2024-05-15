package logger

import (
	"fmt"

	strhelpers "github.com/iolave/go-mach-logger/src/helpers/string"
)

func Info(msg string) {
	fmt.Printf(`{"msg":"%s"}`, strhelpers.ToSnakeCase(msg))
}
