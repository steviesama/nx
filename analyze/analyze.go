// nx/analyze provides a way to analyze various data.
package analyze

import (
  "github.com/steviesama/nx/conv"
)

// HasSliceDuplicates takes an empty interface to a slice and converts it into
// an empty interface slice which it uses to detect duplicates in. If the slice
// param is a slice of primitive data...pass nil to value...otherwise supply a
// function that satisfies ValueFunc.
// It returns a boolean indicating whether duplicates were found.
func HasSliceDuplicates(slice interface{}, value ValueFunc) bool {
  // Tracking return value
  hasDupes := false
  // Convert slice to an empty interface slice
  _slice := conv.InterfaceSlice(slice)

  // A map of each value and the frequency each has been seen in the slice.
  freq := make(map[interface{}]int)

  // Iterate over _slice to map the frequency a given value may occur
  for i := range _slice {
    // primVal represents a primitive datatype for comparison
    var primVal interface{}

    // If the ValueFunc is not nil, the caller wants to supply a primitive
    // datatype for the values in _slice. Or else it will just be treated
    // as a primitive data type.
    if value != nil {
      primVal = value(_slice[i])
    } else {
      primVal = _slice[i]
    }

    // Increment the frequency primVal has been seen.
    freq[primVal]++

    // Check if primVal has been seen more than once
    if freq[primVal] > 1 {
      // Flag that the slice has duplicates.
      hasDupes = true
      // Exit loop
      break
    }
  }

  return hasDupes
}
