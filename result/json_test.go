package result_test

import (
	"encoding/json"
	"fmt"
	"github.com/BooleanCat/go-functional/internal/assert"
	"github.com/BooleanCat/go-functional/result"
	"testing"
)

func ExampleResult_UnmarshalJSON() {
	var r result.Result[[]string]

	_ = json.Unmarshal([]byte(`["Foo", "Bar"]`), &r)

	value, _ := r.Value()

	fmt.Println(value)

	// Output:
	// [Foo Bar]
}

func TestUnmarshalJSON(t *testing.T) {
	var r result.Result[[]string]

	err := json.Unmarshal([]byte(`["Foo", "Bar"]`), &r)
	assert.Nil(t, err)

	value, err := r.Value()
	assert.Nil(t, err)

	assert.SliceEqual(t, value, []string{"Foo", "Bar"})
}

func TestUnmarshalError(t *testing.T) {
	var r result.Result[string]

	err := json.Unmarshal([]byte("42"), &r)

	assert.NotNil(t, err)
}
