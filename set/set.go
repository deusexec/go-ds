package set

type Interface interface {
	Add(...any)
	Sub(...any)
	Xor(*set) *set
	Has(any) bool
	Size() int
	Delete(any)
	Clear()
}

type set struct {
	items map[any]bool
}

func New(items ...any) *set {
	l := len(items)
	s := new(set)
	s.items = make(map[any]bool, l)
	for _, i := range items {
		s.items[i] = true
	}
	return s
}

func (s *set) Add(items ...any) {
	for _, item := range items {
		s.items[item] = true
	}
}

func (s *set) Sub(items ...any) {
	for _, item := range items {
		delete(s.items, item)
	}
}

func (s *set) Xor(s1 *set) *set {
	out := new(set)
	out.items = make(map[any]bool)
	// xor = (s1 - s2) + (s2 - s1)
	for item := range s.items {
		if !s1.Has(item) {
			out.items[item] = true
		}
	}
	for item := range s1.items {
		if !s.Has(item) {
			out.items[item] = true
		}
	}
	return out
}

func (s *set) Has(item any) bool {
	return s.items[item]
}

func (s *set) Delete(item any) {
	delete(s.items, item)
}

func (s *set) Clear() {
	for item := range s.items {
		delete(s.items, item)
	}
}

func (s *set) Size() int {
	return len(s.items)
}
