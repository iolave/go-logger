package strhelpers

import "testing"

func TestToSnakeCase(t *testing.T) {

	type TestCase struct {
		Msg  string
		Want string
	}

	testCases := []TestCase{
		TestCase{"This is a message 123_1 2 4 123", "this_is_a_message_123_1_2_4_123"},
		TestCase{"s1mpl3  m3ssage1", "s_1_mpl_3_m_3_ssage_1"},
		TestCase{"some Weird message 123 testing 1 a 2", "some_weird_message_123_testing_1_a_2"},
	}

	for i := range testCases {
		testCase := testCases[i]
		snakeCased := ToSnakeCase(testCase.Msg)

		if testCase.Want != snakeCased {
			t.Fatalf(`ToSnakeCase("%s"), expected = "%s", got = "%s"`, testCase.Msg, testCase.Want, snakeCased)
		}
	}

}
