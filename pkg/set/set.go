package set;

type Set map[any]struct{}

func (s Set) Has(key any) (bool) {
	_, exists := s[key]
	return exists
}

func (s Set) Add(key any) {
	s[key] = struct{}{}
}