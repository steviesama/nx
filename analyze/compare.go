package analyze

import (
  "strings"
)

// StringCompare is a stock analyze.CompareFunc that a caller can use
// to compare strings without writing an anonymous function.
// It returns an analyze.Equality that represent the state of
// equality between this & that.
func StringCompare(this, that interface{}) Equality {
  // Convert this & that to strings via type assertion.
  // Ignore, the underscore, the value that tells whether
  // the assertion worked or not.
  // TODO: handle the errors that are currently ignored.
  thisStr, _ := this.(string)
  thatStr, _ := that.(string)
  // Run strings.Compare on thisStr/thatStr and convert the result to Equality
  return Equality(strings.Compare(thisStr, thatStr))
}