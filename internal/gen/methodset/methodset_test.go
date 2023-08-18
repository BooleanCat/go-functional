package methodset_test

import (
	"reflect"
	"sort"
	"testing"

	"github.com/BooleanCat/go-functional/internal/assert"
	"github.com/BooleanCat/go-functional/internal/gen/methodset"
	"github.com/BooleanCat/go-functional/iter"
)

func TestAdd(t *testing.T) {
	set := make(methodset.MethodSet)
	set.Add("Drop", "*ChainIter[T]", "T")

	assert.True(t, reflect.DeepEqual(set["Drop"], map[string]string{"*ChainIter[T]": "T"}))
}

func TestMembers(t *testing.T) {
	set := make(methodset.MethodSet)
	set.Add("Drop", "*ChainIter[T]", "T")
	set.Add("Drop", "*ChainIter[T]", "T")
	set.Add("Drop", "*CounterIter", "int")

	members := set.Members("Drop")

	sort.Slice(members, func(i, j int) bool {
		return members[i].One < members[j].One
	})

	assert.SliceEqual(t, members, []iter.Tuple[string, string]{{One: "*ChainIter[T]", Two: "T"}, {One: "*CounterIter", Two: "int"}})
}
