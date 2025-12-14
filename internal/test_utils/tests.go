// go-coverage:ignore
package testutils

import "testing"

func CmpString(t *testing.T, expected, got string) {
	if expected != got {
		t.Fatalf(
			"[TEST_FAILED] %s\n\n\t+ (want) %v\n\t- (got)  %v",
			t.Name(),
			expected,
			got,
		)
	}
}
