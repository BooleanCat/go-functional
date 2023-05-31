package option_test

import (
	"encoding/json"
	"testing"

	"github.com/BooleanCat/go-functional/internal/assert"
	"github.com/BooleanCat/go-functional/option"
)

func TestMarshalSome(t *testing.T) {
	data, err := json.Marshal(option.Some(4))
	assert.Nil(t, err)
	assert.Equal(t, string(data), "4")
}

func TestMarshalNone(t *testing.T) {
	data, err := json.Marshal(option.None[int]())
	assert.Nil(t, err)
	assert.Equal(t, string(data), "null")
}

func TestMarshalSomeParsed(t *testing.T) {
	type name struct {
		MiddleName option.Option[string] `json:"middle_name"`
	}

	data, err := json.Marshal(name{MiddleName: option.Some("Barry")})
	assert.Nil(t, err)
	assert.Equal(t, string(data), `{"middle_name":"Barry"}`)
}

func TestMarshalNoneParsed(t *testing.T) {
	type name struct {
		MiddleName option.Option[string] `json:"middle_name"`
	}

	data, err := json.Marshal(name{MiddleName: option.None[string]()})
	assert.Nil(t, err)
	assert.Equal(t, string(data), `{"middle_name":null}`)
}
