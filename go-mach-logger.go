package logger

import "fmt"

type Logger struct {
	name string
}

func New(name string) Logger {
	logger := new(Logger)

	logger.name = name

	return *logger
}

func (log Logger) Info() {
	fmt.Println("Hello Info")
}

func (log Logger) Warn() {
	fmt.Println("Hello Warn")
}

func (log Logger) Error() {
	fmt.Println("Hello Error")
}

func (log Logger) Debug() {
	fmt.Println("Hello Debug")
}

func (log Logger) Fatal() {
	fmt.Println("Hello Fatal")
}
