package template

import (
	"bytes"
	_ "embed"
	"fmt"
	"text/template"

	"github.com/BooleanCat/go-functional/iter"
	"github.com/BooleanCat/go-functional/result"
)

var (
	//go:embed Header.go.tmpl
	HeaderTemplate []byte

	//go:embed Drop.go.tmpl
	DropTemplate string

	//go:embed Take.go.tmpl
	TakeTemplate string

	//go:embed ForEach.go.tmpl
	ForEachTemplate string
)

var templates = map[string]string{
	"Drop":    DropTemplate,
	"Take":    TakeTemplate,
	"ForEach": ForEachTemplate,
}

type Values struct {
	TypeName   string
	ReturnType string
}

func FromMethodSetItems(items map[string]string) []Values {
	return iter.Map[iter.Tuple[string, string], Values](iter.LiftHashMap(items), func(item iter.Tuple[string, string]) Values {
		return Values{TypeName: item.One, ReturnType: item.Two}
	}).Collect()
}

func RenderTemplate(name string, values []Values) result.Result[[]byte] {
	targetTemplate, ok := templates[name]
	if !ok {
		return result.Err[[]byte](fmt.Errorf("invalid template: %s", name))
	}

	tmpl, _ := template.New(name).Parse(targetTemplate)

	buffer := new(bytes.Buffer)
	_ = tmpl.Execute(buffer, values)

	return result.Ok(buffer.Bytes())
}
