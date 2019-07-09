package value

// Used to assign equality between 2 compared items.
type Equality int

// Values to be returns by CompareFunc functions to indicate whether the items
// being compares are Equal.
// LessThan means this is less than that.
// GreaterThan means this is greater than that.
const (
  LessThan    Equality = -1
  Equal       Equality = 0
  GreaterThan Equality = 1
)

// CompareFunc function signature
type CompareFunc = func(this, that interface{}) Equality
