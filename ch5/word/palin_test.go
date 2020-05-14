package word

import "testing"

func TestIsPalindrome(t *testing.T) {
	if !IsPalindrome("aba") {
		t.Error(`IsPalindrom("aba") = false`)
	}
}

func TestNonPalindrome(t *testing.T) {
	if IsPalindrome("abc") {
		t.Error(`IsPalindrome("abc") = true`)
	}
}

func TestFrenchPalindrome(t *testing.T) {
	if !IsPalindrome("été") {
		t.Error(`IsPalindrome("été") = false`)
	}
}

func TestAll(t *testing.T) {
	var tests = []struct {
		input  string
		expect bool
	}{
		{"aba", true},
		{"abc", false},
		{"", true},
		{"a", true},
		{"kayak", true},
		{"Kayak", true},
		{"été", true},
	}

	for _, test := range tests {
		if res := IsPalindrome(test.input); res != test.expect {
			t.Errorf("IsPalindrome(%q) = %v", test.input, res)
		}
	}
}
