package directive

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/BooleanCat/go-functional/result"
)

var pattern = regexp.MustCompile(`^\/\/gofunctional:generate (?P<Type>\*[A-Z][A-Za-z]*(\[[A-Za-z]+\])?) (?P<YieldedType>[A-Za-z]+)(?P<Methods>( [A-Z][A-Za-z]*)+)$`)

type Directive struct {
	Type        string
	YieldedType string
	Methods     []string
}

func FromString(line string) result.Result[Directive] {
	if !pattern.MatchString(line) {
		return result.Err[Directive](fmt.Errorf("invalid directive: %s", line))
	}

	matches := pattern.FindStringSubmatch(line)
	typeIndex := pattern.SubexpIndex("Type")
	methodsIndex := pattern.SubexpIndex("Methods")
	yieldedTypeIndex := pattern.SubexpIndex("YieldedType")

	return result.Ok(Directive{
		Type:        matches[typeIndex],
		Methods:     strings.Split(strings.TrimSpace(matches[methodsIndex]), " "),
		YieldedType: matches[yieldedTypeIndex],
	})
}
