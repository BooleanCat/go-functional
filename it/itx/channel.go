package itx

import (
	"iter"

	"github.com/BooleanCat/go-functional/v2/it"
)

// ToChannel is a convenience method for chaining [it.ToChannel] on
// [Iterator]s.
func (iterator Iterator[V]) ToChannel() <-chan V {
	return it.ToChannel(iter.Seq[V](iterator))
}
