package strutils

import (
	"testing"
)

func TestToSnakeCase(t *testing.T) {

	type TestCase struct {
		Msg  string
		Want string
	}

	testCases := []TestCase{
		// Empty string
		{"", ""},
		// Single number
		{"1", "1"},
		// Single letter
		{"A", "a"},
		// Single symbol
		{" ", ""},
		// Starts with symbol
		{" test", "test"},
		// Starts with letter
		{"test", "test"},
		// Starts with number
		{"1test", "1_test"},
		// Ends with number
		{"test1", "test_1"},
		// Ends with symbol
		{"test_", "test"},
		// Ends with upper case
		{"tesT", "tes_t"},
		// Ends with 2 symbols
		{"test__", "test"},
		// Ends with char and prev is number
		{"tes1t", "tes_1_t"},
		{"This is a message 123_1 2 4 123", "this_is_a_message_123_1_2_4_123"},
		{"s1mpl3  m_3Ssage1", "s_1_mpl_3_m_3_s_sage_1"},
		{"s1mpl3  m3ssage1", "s_1_mpl_3_m_3_ssage_1"},
		{"some Weird message 123 testing 1 a 2", "some_w_eird_message_123_testing_1_a_2"},
	}

	for i := range testCases {
		testCase := testCases[i]
		snakeCased := ToSnakeCase(testCase.Msg)

		if testCase.Want != snakeCased {
			t.Fatalf(`ToSnakeCase("%s"), expected = "%s", got = "%s"`, testCase.Msg, testCase.Want, snakeCased)
		}
	}
}
