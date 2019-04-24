package models_test

import (
	"testing"

	"github.com/yashdalfthegray/color-info/models"
)

func TestNewErrorResponse(t *testing.T) {
	testCases := []struct {
		desc string
		in   string
		out  models.ErrorResponse
	}{
		{
			desc: "creates a new error response",
			in:   "invalid color",
			out:  models.ErrorResponse{"error", "invalid color"},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			if models.NewErrorResponse(tC.in) != tC.out {
				t.Errorf("expected %s but got %s", tC.out.Message, tC.in)
			}
		})
	}
}
