package option

import "encoding/json"

// MarshalJSON implements the [json.Marshaler] interface.
//
//   - [Some] variants will be marshaled as their underlying value.
//   - [None] variants will be marshaled as "null".
func (o Option[T]) MarshalJSON() ([]byte, error) {
	if value, ok := o.Value(); ok {
		return json.Marshal(value)
	}

	return []byte("null"), nil
}

// UnmarshalJSON implements the [json.Unmarshaler] interface.
//
//   - Values will be marshed as [Some] variants.
//   - "null"s will be marshaled as [None] variants.
func (o *Option[T]) UnmarshalJSON(data []byte) error {
	*o = None[T]()

	if string(data) != "null" {
		var value T
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		*o = Some(value)
	}

	return nil
}

var (
	_ json.Unmarshaler = &Option[struct{}]{}
	_ json.Marshaler   = Option[struct{}]{}
)
