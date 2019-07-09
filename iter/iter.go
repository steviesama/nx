// nx/iter is an iterator package that provides interfaces that various user-defined
// types can use to become iterable.
package iter

import (
)

// iter.Traversable is an interface for iterating over anything that implements this
// interface. Unlike a normal Iterator, a Traverser
type Traverser interface {  
  // Prev return true if there is another element in going backward.
  Prev() bool
  // Next returns true if there is another element going forward.
  Next() bool
  // Dec, decrements the index backward.
  Dec()
  // Inc, increments the index forward.
  Inc()
  // Value returns the value
  Value() interface{}
}

// Traversable is used to hold the traverable object
type Traversable struct {  
  Traverser
  // iterant is the original object converted to an empty interface slice
  iterant []interface{}
  index int
}

func (tr *Traversable) Next() bool {
  return true
}