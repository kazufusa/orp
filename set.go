package orp

type Set = []string

func NewSet(ss []string) Set {
	m := make(map[string]struct{}, len(ss))
	for _, s := range ss {
		if s == "" {
			continue
		}
		m[s] = struct{}{}
	}
	ret := make(Set, 0, len(m))
	for key := range m {
		ret = append(ret, key)
	}
	return ret
}
