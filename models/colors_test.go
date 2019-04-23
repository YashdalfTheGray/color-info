package models_test

import (
	"testing"

	"github.com/yashdalfthegray/color-info/models"
)

func TestHexCodeToRGB(t *testing.T) {
	testCases := []struct {
		desc          string
		in            string
		out           string
		errorExpected bool
	}{
		{
			desc:          "converts valid color string",
			in:            "#164080",
			out:           "rgb(22, 64, 128)",
			errorExpected: false,
		},
		{
			desc:          "converts black color string",
			in:            "#000000",
			out:           "rgb(0, 0, 0)",
			errorExpected: false,
		},
		{
			desc:          "converts white color string",
			in:            "#ffffff",
			out:           "rgb(255, 255, 255)",
			errorExpected: false,
		},
		{
			desc:          "errors for invalid color string",
			in:            "#1640",
			out:           "rgb(0, 0, 0)",
			errorExpected: true,
		},
		{
			desc:          "errors for empty string",
			in:            "",
			out:           "rgb(0, 0, 0)",
			errorExpected: true,
		},
		{
			desc:          "errors for malformed string",
			in:            "164080",
			out:           "rgb(0, 0, 0)",
			errorExpected: true,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			result, err := models.HexCode(tC.in).ToRGB()
			if !tC.errorExpected && err != nil {
				t.Errorf(err.Error())
			}
			if tC.errorExpected && err == nil {
				t.Errorf("expecting error to not be nil")
			}
			if result.String() != tC.out {
				t.Errorf("expected %s but got %s", tC.out, result.String())
			}
		})
	}
}

func TestRGBString(t *testing.T) {
	testCases := []struct {
		desc string
		in   models.RGB
		out  string
	}{
		{
			desc: "prints out a color in standard RGB format",
			in:   models.RGB{22, 64, 128},
			out:  "rgb(22, 64, 128)",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			if tC.in.String() != tC.out {
				t.Errorf("expected %s but got %s", tC.out, tC.in.String())
			}
		})
	}
}
