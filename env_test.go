package orp

import (
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestEnv(t *testing.T) {
	given := "/a:/a/b/c:/a/b::/c:/d/e/f:/d/e:/d/e/f"
	name := "TESTPATH"
	err := os.Setenv(name, given)
	if err != nil {
		t.Errorf("Error should not be occured: %s", err)
	}
	e := NewEnv(name, ":")
	e.MoveToTop("^/a/b")
	e.MoveToTop("e")
	expected := "/d/e/f:/d/e:/a/b/c:/a/b:/a:/c"
	err = e.Export()
	if err != nil {
		t.Errorf("Error should not be occured: %s", err)
	}
	actual := os.Getenv("TESTPATH")
	if diff := cmp.Diff(expected, actual); diff != "" {
		t.Errorf("given(%s): expected %s, actual %s\n%s", given, expected, actual, diff)
	}
}

func TestNewEnv(t *testing.T) {
	opt := cmpopts.SortSlices(func(i, j string) bool {
		return i < j
	})

	given := "A:B:C:D:E:::"
	expected := []string{"A", "B", "C", "D", "E"}

	err := os.Setenv("TESTPATH", given)
	if err != nil {
		t.Errorf("Error should not be occured: %s", err)
	}
	e := NewEnv("TESTPATH", ":")
	actual := e.items
	if diff := cmp.Diff(expected, actual, opt); diff != "" {
		t.Errorf("given(%s): expected %s, actual %s\n%s", given, expected, actual, diff)
	}
}

func TestEnvMoveToTop(t *testing.T) {
	var tests = []struct {
		name     string
		arg      []string
		expected []string
		given    []string
	}{
		{
			"simple",
			[]string{"1", "2", "3", "4"},
			[]string{"4", "2", "1", "3"},
			[]string{"4", "2"},
		},
		{
			"regexp",
			[]string{"/a", "/b/c", "/b/c/d", "/e", "/b/b/b/c", "/b"},
			[]string{"/b/c", "/b/c/d", "/e", "/a", "/b/b/b/c", "/b"},
			[]string{"^/b/c", "e"},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			e := Env{items: tt.arg}
			for i := len(tt.given) - 1; i >= 0; i-- {
				e.MoveToTop(tt.given[i])
			}
			actual := e.items
			if !cmp.Equal(tt.expected, actual) {
				t.Errorf("given(%s): expected %s, actual %s\n%s", tt.given, tt.expected, actual, cmp.Diff(tt.expected, actual))
			}
		})
	}
}
