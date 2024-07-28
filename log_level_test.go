package logger

import (
	"os"
	"testing"
)

func TestLogLevelShouldBeInfo(t *testing.T) {
	want := "info"

	type testCase struct {
		Level LogLevel
		Want  string
	}

	testCases := []testCase{
		{Level: LOG_LEVEL_INFO + 1, Want: want},
		{Level: LOG_LEVEL_INFO, Want: want},
	}

	for i := range testCases {
		testCase := testCases[i]
		result := testCase.Level.ToString()

		t.Logf(`LogLevel.ToString("%d") = %s`, testCase.Level, result)

		if testCase.Want != result {
			t.Fatalf(`LogLevel.ToString("%d") = %s, want = "%s", got = "%s"`, testCase.Level, result, testCase.Want, result)
		}
	}
}

func TestLogLevelShouldBeDebug(t *testing.T) {
	want := "debug"

	type testCase struct {
		Level LogLevel
		Want  string
	}

	testCases := []testCase{
		{Level: LOG_LEVEL_DEBUG + 1, Want: want},
		{Level: LOG_LEVEL_DEBUG, Want: want},
	}

	for i := range testCases {
		testCase := testCases[i]
		result := testCase.Level.ToString()

		t.Logf(`LogLevel.ToString("%d") = %s`, testCase.Level, result)

		if testCase.Want != result {
			t.Fatalf(`LogLevel.ToString("%d") = %s, want = "%s", got = "%s"`, testCase.Level, result, testCase.Want, result)
		}
	}
}

func TestLogLevelShouldBeNotSet(t *testing.T) {
	want := "not_set"

	type testCase struct {
		Level LogLevel
		Want  string
	}

	testCases := []testCase{
		{Level: LOG_LEVEL_DEBUG - 1, Want: want},
		{Level: LOG_LEVEL_NOTSET + 1, Want: want},
		{Level: LOG_LEVEL_NOTSET, Want: want},
	}

	for i := range testCases {
		testCase := testCases[i]
		result := testCase.Level.ToString()

		t.Logf(`LogLevel.ToString("%d") = %s`, testCase.Level, result)

		if testCase.Want != result {
			t.Fatalf(`LogLevel.ToString("%d") = %s, want = "%s", got = "%s"`, testCase.Level, result, testCase.Want, result)
		}
	}
}
func TestLogLevelShouldBeError(t *testing.T) {
	want := "error"

	type testCase struct {
		Level LogLevel
		Want  string
	}

	testCases := []testCase{
		{Level: LOG_LEVEL_ERROR + 1, Want: want},
		{Level: LOG_LEVEL_ERROR, Want: want},
	}

	for i := range testCases {
		testCase := testCases[i]
		result := testCase.Level.ToString()

		t.Logf(`LogLevel.ToString("%d") = %s`, testCase.Level, result)

		if testCase.Want != result {
			t.Fatalf(`LogLevel.ToString("%d") = %s, want = "%s", got = "%s"`, testCase.Level, result, testCase.Want, result)
		}
	}
}
func TestLogLevelShouldBeWarn(t *testing.T) {
	want := "warn"

	type testCase struct {
		Level LogLevel
		Want  string
	}

	testCases := []testCase{
		{Level: LOG_LEVEL_WARN + 1, Want: want},
		{Level: LOG_LEVEL_WARN, Want: want},
	}

	for i := range testCases {
		testCase := testCases[i]
		result := testCase.Level.ToString()

		t.Logf(`LogLevel.ToString("%d") = %s`, testCase.Level, result)

		if testCase.Want != result {
			t.Fatalf(`LogLevel.ToString("%d") = %s, want = "%s", got = "%s"`, testCase.Level, result, testCase.Want, result)
		}
	}
}
func TestLogLevelShouldBeFatal(t *testing.T) {
	want := "fatal"

	type testCase struct {
		Level LogLevel
		Want  string
	}

	testCases := []testCase{
		{Level: LOG_LEVEL_FATAL + 1, Want: want},
		{Level: LOG_LEVEL_FATAL, Want: want},
	}

	for i := range testCases {
		testCase := testCases[i]
		result := testCase.Level.ToString()

		t.Logf(`LogLevel.ToString("%d") = %s`, testCase.Level, result)

		if testCase.Want != result {
			t.Fatalf(`LogLevel.ToString("%d") = %s, want = "%s", got = "%s"`, testCase.Level, result, testCase.Want, result)
		}
	}
}

func TestGetLogLevelFromEnv(t *testing.T) {
	type testCase struct {
		Want LogLevel
		Env  string
	}

	testCases := []testCase{
		{Want: LOG_LEVEL_FATAL, Env: "Fatal"},
		{Want: LOG_LEVEL_ERROR, Env: "eRRor"},
		{Want: LOG_LEVEL_WARN, Env: "warn"},
		{Want: LOG_LEVEL_INFO, Env: "info"},
		{Want: LOG_LEVEL_DEBUG, Env: "debug"},
		{Want: LOG_LEVEL_INFO, Env: "something_else"},
	}

	for i := range testCases {
		testCase := testCases[i]
		os.Setenv("LOG_LEVEL", testCase.Env)

		result := getLogLevelFromEnv()

		t.Logf(`getLogLevelFromEnv() = %d`, result)

		if testCase.Want != result {
			t.Fatalf(`getLogLevelFromEnv() = %d, want = "%d", got = "%d"`, testCase.Want, result, testCase.Want)
		}
	}
}
