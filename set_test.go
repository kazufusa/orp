package orp

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestSet(t *testing.T) {
	opt := cmpopts.SortSlices(func(i, j string) bool {
		return i < j
	})
	var tests = []struct {
		expected []string
		given    []string
	}{
		{
			[]string{"1", "2", "3"},
			[]string{"1", "2", "3", ""},
		},
		{
			[]string{"1", "2", "3", "4", "10"},
			[]string{"1", "2", "3", "2", "3", "", "4", "10"},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(fmt.Sprintf("%v", tt.given), func(t *testing.T) {
			actual := NewSet(tt.given)
			if diff := cmp.Diff(tt.expected, actual, opt); diff != "" {
				t.Errorf("given(%s): expected %s, actual %s\n%s", tt.given, tt.expected, actual, diff)
			}
		})
	}
}
