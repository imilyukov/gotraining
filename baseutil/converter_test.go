package baseutil

import (
	"testing"
	"fmt"
)

func TestNumb10_BaseOf(t *testing.T) {
	cases := []struct {
		base byte
		number10 int64
		result BasedNumb
	}{
		{35, 1365333188, BasedNumb{Base: 35, Numb: "python"}},
		{33, 34297654,   BasedNumb{Base: 33, Numb: "sucks"}},
		{33, 37,         BasedNumb{Base: 33, Numb: "14"}},
		{2,  127,        BasedNumb{Base: 2,  Numb: "1111111"}},
	}

	for _, c := range cases {
		var g, r BasedNumb = Numb10(c.number10).BaseOf(c.base), c.result
		if g != r {
			t.Errorf(
				"Numb10(%d).BaseOf(%d) == BasedNumb{Base: %d, Numb: %s}, want BasedNumb{Base: %d, Numb: %s}",
				c.number10,
				c.base,
				g.Base,
				g.Numb,
				r.Base,
				r.Numb,
			)
		}
	}
}

func TestANSI_IndexOf(t *testing.T) {
	cases := []struct {
		input  byte
		result string
	}{
		{0,  "0"},
		{1,  "1"},
		{2,  "2"},
		{3,  "3"},
		{4,  "4"},
		{5,  "5"},
		{6,  "6"},
		{7,  "7"},
		{8,  "8"},
		{9,  "9"},
		{10, "a"},
		{11, "b"},
		{12, "c"},
		{13, "d"},
		{14, "e"},
		{15, "f"},
	}

	for _, c := range cases {
		var g, r string = string(ANSI(c.input).IndexOf()), c.result
		fmt.Print(ANSI(c.input).IndexOf())
		if g != r {
			t.Errorf(
				"ANSI(%d).IndexOf() == %s, want %s",
				c.input,
				g,
				r,
			)
		}
	}
}