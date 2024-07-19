package itx

import "github.com/BooleanCat/go-functional/v2/it"

// Left is a convenience method that unzips an [Iterator2] and returns the left
// iterator, closing the right iterator.
func (iterator Iterator2[V, W]) Left() Iterator[V] {
	return Iterator[V](it.Left(iterator))
}

// Right is a convenience method that unzips an [Iterator2] and returns the
// right iterator, closing the left iterator.
func (iterator Iterator2[V, W]) Right() Iterator[W] {
	return Iterator[W](it.Right(iterator))
}
