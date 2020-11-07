package ArraysAndStrings

import (
	"testing"
)

func TestIsUnique(t *testing.T) {
	testCases := []struct {
		input  string
		expect bool
	}{
		{"abc", true},
		{"aab", false},
		{"abcdefghijklmnopqrstuvxywz`1234567890-=", true},
		{"abcdefghijklmnopqrstuvxywz`1234567890-=~!@#$%^&*()_+", true},
		{"abcdefghijklmnopqrstuvxywz`1234567890-=~!@#$%^&*()_+", true},
		{"abcdefghijklmnopqrstuvxywzABCDEFGHIJKLMNOPQRSTUVXYWZ`1234567890-=~!@#$%^&*()_+", true},
		{"abcdefghijklmnopqrstuvxywzABCDEFGHIJKLMNOPQRSTUVXYWZ`1234567890-=~!@#$%^&*()_+,./<>?;'[]{} ", true},
	}

	for _, testCase := range testCases {
		output := IsUnique(testCase.input)
		if output != testCase.expect {
			t.Errorf("Want: %t, Have: %t", testCase.expect, output)
		}
	}
}
