package models_test

import (
	"testing"

	"github.com/yashdalfthegray/color-info/models"
)

func TestNewStatusResponse(t *testing.T) {
	testCases := []struct {
		desc string
		in   string
		out  models.StatusResponse
	}{
		{
			desc: "creates a new error response",
			in:   "test_server_id",
			out:  models.StatusResponse{"ok", "test_server_id"},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			if models.NewStatusResponse(tC.in) != tC.out {
				t.Errorf("expected %s but got %s", tC.out.ServerID, tC.in)
			}
		})
	}
}
