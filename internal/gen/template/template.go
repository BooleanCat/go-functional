package template

import "github.com/BooleanCat/go-functional/iter"

type Values struct {
	TypeName   string
	ReturnType string
}

func FromMethodSetItems(items map[string]string) []Values {
	return iter.Map[iter.Tuple[string, string], Values](iter.LiftHashMap(items), func(item iter.Tuple[string, string]) Values {
		return Values{TypeName: item.One, ReturnType: item.Two}
	}).Collect()
}
