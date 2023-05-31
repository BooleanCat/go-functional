package option

import "encoding/json"

func (o Option[T]) MarshalJSON() ([]byte, error) {
	if value, ok := o.Value(); ok {
		return json.Marshal(value)
	}

	return []byte("null"), nil
}

var _ json.Marshaler = Option[struct{}]{}

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

var _ json.Unmarshaler = &Option[struct{}]{}
