package orp

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

type Env struct {
	items        []string
	name, sep, s string
}

func NewEnv(name, sep string) *Env {
	s := os.Getenv(name)
	return &Env{name: name, sep: sep, s: s, items: NewSet(strings.Split(s, sep))}
}

func (e *Env) Export() string {
	return fmt.Sprintf("%s=%s", e.name, strings.Join(e.items, e.sep))
}

func (e *Env) MoveToTop(s string) error {
	newItems := make([]string, 0, len(e.items))
	re, err := regexp.Compile(s)
	if err != nil {
		return err
	}
	for i := len(e.items) - 1; i >= 0; i-- {
		if re.MatchString(e.items[i]) {
			newItems = append([]string{e.items[i]}, newItems...)
			e.items = append(e.items[:i], e.items[i+1:]...)
		}
	}
	e.items = append(newItems, e.items...)
	return nil
}
