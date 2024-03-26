package main

import (
	"fmt"

	strhelpers "iolave.com/go-app-factory/src/helpers/string"
)

func main() {
	fmt.Println(strhelpers.ToSnakeCase("$SomeLog message v123 to! be_Converted_123"))
}
