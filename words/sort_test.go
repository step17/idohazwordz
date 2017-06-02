package words

import "testing"

func TestSort(t *testing.T) {
	for _, tc := range []struct {
		in   string
		want string
	}{
		{"", ""},
		{"ABC", "ABC"},
		{"CAT", "ACT"},
		{"cat", "act"},
		{"👩⌘💻", "⌘👩💻"},
	} {
		if got := Sort(tc.in); got != tc.want {
			t.Errorf("Sort(%q) = %q want %q", tc.in, got, tc.want)
		}
	}
}
