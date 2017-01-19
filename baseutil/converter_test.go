package baseutil

import "testing"

func TestConvert(t *testing.T) {
	cases := []struct {
		base, number10, result string
	}{
		{"35", "1365333188", "python"},
		{"33", "34297654",   "sucks"},
		{"33", "37",         "14"},
		{"2",  "127",        "1111111"},
	}

	for _, c := range cases {
		got := Convert(c.base, c.number10)
		if got != c.result {
			t.Errorf("Convert(%q, %q) == %q, want %q", c.base, c.number10, got, c.result)
		}
	}
}