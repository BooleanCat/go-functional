package itx

import "github.com/BooleanCat/go-functional/v2/it"

// FromChannel yields values from a channel.
//
// In order to avoid a deadlock, the channel must be closed before attempting
// to called `stop` on a pull-style iterator.
func FromChannel[V any](channel <-chan V) Iterator[V] {
	return Iterator[V](it.FromChannel(channel))
}

// ToChannel is a convenience method for chaining [it.ToChannel] on
// [Iterator]s.
func (iterator Iterator[V]) ToChannel() <-chan V {
	return it.ToChannel(iterator)
}
