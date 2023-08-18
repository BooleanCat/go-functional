package template_test

import (
	"sort"
	"testing"

	"github.com/BooleanCat/go-functional/internal/assert"
	"github.com/BooleanCat/go-functional/internal/gen/template"
)

func TestFromMethodSetItems(t *testing.T) {
	items := make(map[string]string)
	items["*ChainIter[T]"] = "T"
	items["*CountIter"] = "int"

	values := template.FromMethodSetItems(items)

	sort.Slice(values, func(i, j int) bool {
		return values[i].TypeName < values[j].TypeName
	})

	assert.SliceEqual(t, values, []template.Values{
		{TypeName: "*ChainIter[T]", ReturnType: "T"},
		{TypeName: "*CountIter", ReturnType: "int"},
	})
}
