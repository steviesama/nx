package analyze

// Used to assign equality between 2 compared items.
type Equality int

// Values to be returned by CompareFunc functions to indicate whether the items
// being compared are Equal.
// LessThan means this is less than that.
// GreaterThan means this is greater than that.
const (
  LessThan    Equality = -1
  Equal       Equality = 0
  GreaterThan Equality = 1
)

// CompareFunc functions are sent 2 elements and the anonymous function
// referred to by CompareFunc is responsible for determining the Equality
// of this & that and return it.
type CompareFunc = func(this, that interface{}) Equality
