package gollection

type Gollectable any

type callback[T Gollectable] func(value *T)

type Gollection[T Gollectable] struct {
	gollection []*T
	has        map[*T]struct{}
}

func New[T Gollectable]() *Gollection[T] {
	return &Gollection[T]{
		gollection: make([]*T, 0, 100),
		has:        map[*T]struct{}{},
	}
}

func (g *Gollection[T]) Has(target *T) bool {
	_, has := g.has[target]
	return has
}

func (g *Gollection[T]) Add(target *T) {
	if _, has := g.has[target]; has == false {
		g.has[target] = struct{}{}
		g.gollection = append(g.gollection, target)
	}
}

func (g *Gollection[T]) Size() int {
	return len(g.has)
}

func (g *Gollection[T]) Remove(target *T) {
	if _, has := g.has[target]; has {
		delete(g.has, target)
		for i, v := range g.gollection {
			if v == target {
				arr := g.gollection

				switch {
				case len(arr) == 1:
					arr = make([]*T, 0)
				default:
					// swap last element with element to be removed
					arr[i] = arr[len(arr)-1]
					fallthrough
				case len(arr)-1 == i: // last element
					arr = arr[:len(arr)-1]
				}

				g.gollection = arr
			}
		}
	}
}

func (g *Gollection[T]) ForEach(cb callback[T]) {
	for _, v := range g.gollection {
		cb(v)
	}
}

func (g *Gollection[T]) GoEach(cb callback[T]) {
	for _, v := range g.gollection {
		go cb(v)
	}
}
