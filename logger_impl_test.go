package logger

import (
	"context"
	"os"
	"os/exec"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	t.Run("should return an error if the level is not valid", func(t *testing.T) {
		_, err := New(0, "", "")
		if err == nil {
			t.Errorf("expected error, got nil")
		}
	})
	t.Run("should return a logger", func(t *testing.T) {
		_, err := New(LEVEL_INFO, "", "")
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
	})
}

func Test_defaultLogger_log(t *testing.T) {
	t.Run("should not log if the level is lower than the logger level", func(t *testing.T) {
		l := defaultLogger{level: 999999}
		l.log(Entry{Level: LEVEL_TRACE})
	})

	t.Run("should log if the level is higher than the logger level", func(t *testing.T) {
		l := defaultLogger{level: 0}
		l.log(Entry{Level: LEVEL_INFO})
	})
}

func Test_defaultLogger_Trace(t *testing.T) {
	t.Run("should log a trace entry", func(t *testing.T) {
		l := defaultLogger{level: 999999}
		l.Trace(context.Background(), "msg")
	})
}

func Test_defaultLogger_TraceWithData(t *testing.T) {
	t.Run("should log a trace entry", func(t *testing.T) {
		l := defaultLogger{level: 999999}
		l.TraceWithData(context.Background(), "msg", map[string]any{})
	})
}

func Test_defaultLogger_Debug(t *testing.T) {
	t.Run("should log a debug entry", func(t *testing.T) {
		l := defaultLogger{level: 999999}
		l.Debug(context.Background(), "msg")
	})
}

func Test_defaultLogger_DebugWithData(t *testing.T) {
	t.Run("should log a debug entry", func(t *testing.T) {
		l := defaultLogger{level: 999999}
		l.DebugWithData(context.Background(), "msg", map[string]any{})
	})
}

func Test_defaultLogger_Info(t *testing.T) {
	t.Run("should log an info entry", func(t *testing.T) {
		l := defaultLogger{level: 999999}
		l.Info(context.Background(), "msg")
	})
}

func Test_defaultLogger_InfoWithData(t *testing.T) {
	t.Run("should log an info entry", func(t *testing.T) {
		l := defaultLogger{level: 999999}
		l.InfoWithData(context.Background(), "msg", map[string]any{})
	})
}

func Test_defaultLogger_Warn(t *testing.T) {
	t.Run("should log a warn entry", func(t *testing.T) {
		l := defaultLogger{level: 999999}
		l.Warn(context.Background(), "msg", nil)
	})
}

func Test_defaultLogger_WarnWithData(t *testing.T) {
	t.Run("should log a warn entry", func(t *testing.T) {
		l := defaultLogger{level: 999999}
		l.WarnWithData(context.Background(), "msg", nil, map[string]any{})
	})
}

func Test_defaultLogger_Error(t *testing.T) {
	t.Run("should log an error entry", func(t *testing.T) {
		l := defaultLogger{level: 999999}
		l.Error(context.Background(), "msg", nil)
	})
}

func Test_defaultLogger_ErrorWithData(t *testing.T) {
	t.Run("should log an error entry", func(t *testing.T) {
		l := defaultLogger{level: 999999}
		l.ErrorWithData(context.Background(), "msg", nil, map[string]any{})
	})
}

func Test_defaultLogger_Fatal(t *testing.T) {
	t.Run("should log a fatal entry", func(t *testing.T) {
		// Run the crashing code when FLAG is set
		if os.Getenv("FLAG") == "1" {
			l := defaultLogger{level: 999999}
			l.Fatal(context.Background(), "msg", nil)
			return
		}
		// Run the test in a subprocess
		cmd := exec.Command(os.Args[0], "-test.run=Test_defaultLogger_Fatal")
		cmd.Env = append(os.Environ(), "FLAG=1")
		err := cmd.Run()

		// Cast the error as *exec.ExitError and compare the result
		e, ok := err.(*exec.ExitError)
		expectedErrorString := "exit status 1"
		assert.Equal(t, true, ok)
		assert.Equal(t, expectedErrorString, e.Error())
	})
}

func Test_defaultLogger_FatalWithData(t *testing.T) {
	t.Run("should log a fatal entry", func(t *testing.T) {
		// Run the crashing code when FLAG is set
		if os.Getenv("FLAG") == "1" {
			l := defaultLogger{level: 999999}
			l.FatalWithData(context.Background(), "msg", nil, map[string]any{})
			return
		}
		// Run the test in a subprocess
		cmd := exec.Command(os.Args[0], "-test.run=Test_defaultLogger_FatalWithData")
		cmd.Env = append(os.Environ(), "FLAG=1")
		err := cmd.Run()

		// Cast the error as *exec.ExitError and compare the result
		e, ok := err.(*exec.ExitError)
		expectedErrorString := "exit status 1"
		assert.Equal(t, true, ok)
		assert.Equal(t, expectedErrorString, e.Error())
	})
}
