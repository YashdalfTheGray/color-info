package utils_test

import (
	"testing"

	"github.com/yashdalfthegray/color-info/utils"
)

func TestValidateColorString(t *testing.T) {
	testCases := []struct {
		desc string
		in   string
		out  bool
	}{
		{
			desc: "with valid color string",
			in:   "#ffffff",
			out:  true,
		},
		{
			desc: "with another valid color string",
			in:   "#123456",
			out:  true,
		},
		{
			desc: "with yet another valid color string",
			in:   "#b16a12",
			out:  true,
		},
		{
			desc: "with a color string without '#'",
			in:   "3ca12b",
			out:  false,
		},
		{
			desc: "with empty string",
			in:   "",
			out:  false,
		},
		{
			desc: "with more than 3 hex numbers",
			in:   "#fffffff",
			out:  false,
		},
		{
			desc: "with gibberish",
			in:   "gibberish",
			out:  false,
		},
		{
			desc: "with symbols",
			in:   "(*&#(#(#*&",
			out:  false,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			if result := utils.ValidateColorString(tC.in); result != tC.out {
				t.Errorf("expected %t but got %t", tC.out, result)
			}
		})
	}
}
