// nx/iter is an Interface package that provides interfaces that various user-defined
// types can use to become iterable.
// 
// NOTE: Not sure this package is really needed. Development on hold until it proves
//       itself needed.
package iter

// Interface is an interface for objects that want to capable of being iterated over.
type Interface interface {  
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
