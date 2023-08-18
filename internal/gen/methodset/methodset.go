package methodset

import "github.com/BooleanCat/go-functional/iter"

type MethodSet map[string]map[string]string

func (set MethodSet) Add(method, typeName, yieldedType string) {
	if _, ok := set[method]; !ok {
		set[method] = make(map[string]string)
	}

	set[method][typeName] = yieldedType
}

func (set MethodSet) Members(method string) []iter.Tuple[string, string] {
	return iter.Collect[iter.Tuple[string, string]](iter.LiftHashMap(set[method]))
}
