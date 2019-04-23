package utils_test

import (
	"testing"

	"github.com/yashdalfthegray/color-info/utils"
)

func TestGetServerID(t *testing.T) {
	testCases := []struct {
		desc string
	}{
		{
			desc: "Gets a new server ID every invocation",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			s1 := utils.GetServerID()
			s2 := utils.GetServerID()

			if s1 == s2 {
				t.Errorf("GetServerID returned the same ID")
			}
		})
	}
}
