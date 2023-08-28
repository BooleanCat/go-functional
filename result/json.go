package result

import "encoding/json"

// UnmarshalJSON implements the [json.Unmarshaler] interface.
// Values will be unmarshaled as [Ok] variants.
func (r *Result[T]) UnmarshalJSON(data []byte) error {
	var value T

	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}

	*r = Ok(value)
	return nil
}
