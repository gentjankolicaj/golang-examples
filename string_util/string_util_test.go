package string_util

import "testing"

func TestReverseRunes(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"Hello, world", "dlrow ,olleH"},
		{"Hello, 世界", "界世 ,olleH"},
		{"", ""},
	}

	for _, v := range cases {
		reverseStr := ReverseRunes(v.in)

		if reverseStr != v.want {
			t.Errorf("ReverseRunes(%q) == %q, want %q", v.in, reverseStr, v.want)
		}
	}
}
