package orp

type Set = []string

func NewSet(ss []string) Set {
	ret := make(Set, 0, len(ss))
	m := make(map[string]struct{}, len(ss))
	for _, s := range ss {
		if s == "" {
			continue
		}
		if _, ok := m[s]; ok {
			continue
		}
		ret = append(ret, s)
		m[s] = struct{}{}
	}
	return ret
}
