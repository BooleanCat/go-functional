package option

import "encoding/json"

func (o Option[T]) MarshalJSON() ([]byte, error) {
	if value, ok := o.Value(); ok {
		return json.Marshal(value)
	}

	return []byte("null"), nil
}

var _ json.Marshaler = Option[struct{}]{}
