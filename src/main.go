package main

import (
	"fmt"
	"os"

	strhelpers "iolave.com/go-app-factory/src/helpers/string"
)

func main() {
	fmt.Println(strhelpers.ToSnakeCase("SomeLog message to be_Converted"))

	os.Exit(0)
}
