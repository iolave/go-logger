package logger

import (
	"os"
	"os/exec"
	"testing"
)

func TestAppShouldExitWithCode1(t *testing.T) {
	if os.Getenv("BE_CRASHER") == "1" {
		logger := New("test")
		logger.Fatal("msg", map[string]any{})
		return
	}

	cmd := exec.Command(os.Args[0], "-test.run=TestAppShouldExitWithCode1")
	cmd.Env = append(os.Environ(), "BE_CRASHER=1")
	err := cmd.Run()
	if e, ok := err.(*exec.ExitError); ok && !e.Success() {
		return
	}

	t.Fatalf(`Logger.Fatal(), should exit with code = 1, got = %d`, err)
}
