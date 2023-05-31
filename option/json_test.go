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

func TestUnmarshalSome(t *testing.T) {
	var number option.Option[int]
	err := json.Unmarshal([]byte("4"), &number)
	assert.Nil(t, err)
	assert.Equal(t, number, option.Some(4))
}

func TestUnmarshalNone(t *testing.T) {
	var number option.Option[int]
	err := json.Unmarshal([]byte("null"), &number)
	assert.Nil(t, err)
	assert.True(t, number.IsNone())
}

func TestUnmarshalEmpty(t *testing.T) {
	type name struct {
		MiddleName option.Option[string] `json:"middle_name"`
	}

	var value name
	err := json.Unmarshal([]byte("{}"), &value)
	assert.Nil(t, err)
	assert.True(t, value.MiddleName.IsNone())
}
