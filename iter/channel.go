package iter

import "iter"

// FromChannel yields values from a channel.
//
// In order to avoid a deadlock, the channel must be closed before attempting
// to called `stop` on a pull-style iterator.
func FromChannel[V any](channel <-chan V) iter.Seq[V] {
	return func(yield func(V) bool) {
		for value := range channel {
			if !yield(value) {
				return
			}
		}
	}
}

// ToChannel sends yielded values to a channel.
//
// The channel is closed when the iterator is exhausted. Beware of leaked go
// routines when using this function with an infinite iterator.
func ToChannel[V any](seq iter.Seq[V]) <-chan V {
	channel := make(chan V)

	go func() {
		defer close(channel)

		for value := range seq {
			channel <- value
		}
	}()

	return channel
}

// ToChannel is a convenience method for chaining [ToChannel] on [Iterator]s.
func (iterator Iterator[V]) ToChannel() <-chan V {
	return ToChannel(iter.Seq[V](iterator))
}
