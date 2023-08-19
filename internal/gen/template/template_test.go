package template_test

import (
	_ "embed"
	"sort"
	"strings"
	"testing"

	"github.com/BooleanCat/go-functional/internal/assert"
	"github.com/BooleanCat/go-functional/internal/gen/template"
)

//go:embed fixtures/foreach.txt
var renderedForEach string

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

func TestRenderTemplate(t *testing.T) {
	content := template.RenderTemplate("ForEach", []template.Values{{"*CountIter", "int"}}).Unwrap()
	assert.Equal(t, strings.TrimSpace(string(content)), strings.TrimSpace(renderedForEach))
}
