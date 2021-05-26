package morestrings

import "testing"

func TestExclaim(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"hello", "HELLO!!!"},
		{"went boom", "WENT BOOM!!!"},
	}
	for _, c := range cases {
		got := Exclaim(c.in)
		if got != c.want {
			t.Errorf("Exclaim(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}
