package logger

import (
	"os"
	"testing"
)

func TestShouldPrintWhenLogLevelIsTheSame(t *testing.T) {
	os.Setenv("LOG_LEVEL", "info")

	entry := LogEntry{
		Level: LOG_LEVEL_INFO,
	}

	want := true
	result := entry.shouldPrint()
	t.Logf(`LogEntry.shouldPrint() = %t`, result)

	if want != result {
		t.Fatalf(`LogEntry.shouldPrint() = %t, want = "%t", got = "%t"`, result, want, result)
	}

}

func TestShouldNotPrintWhenEnvLogLevelIsHigher(t *testing.T) {
	os.Setenv("LOG_LEVEL", "fatal")

	entry := LogEntry{
		Level: LOG_LEVEL_INFO,
	}

	result := entry.shouldPrint()
	want := false
	t.Logf(`LogEntry.shouldPrint() = %t`, result)

	if want != result {
		t.Fatalf(`LogEntry.shouldPrint() = %t, want = "%t", got = "%t"`, result, want, result)
	}

}

func TestShouldPrintWhenEnvLogLevelIsLower(t *testing.T) {
	os.Setenv("LOG_LEVEL", "debug")

	entry := LogEntry{
		Level: LOG_LEVEL_INFO,
	}

	result := entry.shouldPrint()
	want := true
	t.Logf(`LogEntry.shouldPrint() = %t`, result)

	if want != result {
		t.Fatalf(`LogEntry.shouldPrint() = %t, want = "%t", got = "%t"`, result, want, result)
	}

}
