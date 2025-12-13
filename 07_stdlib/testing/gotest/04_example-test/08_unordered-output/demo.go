package demo

type Set map[string]struct{}

func (s Set) Add(v string) {
	s[v] = struct{}{}
}

func (s Set) Values() []string {
	var out []string
	for v := range s {
		out = append(out, v)
	}
	return out
}
