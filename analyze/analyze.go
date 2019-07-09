// nx/analyze provides a way to analyze various data.
package analyze

import (
  "github.com/steviesama/nx/conv"
)

func HasSliceDuplicates(slice interface{}, compare CompareFunc) bool {
  // Tracking return value
  hasDupes := false
  // Convert slice to an empty interface slice
  _slice := conv.InterfaceSlice(slice)

  // A map of each value and the frequency each has been seen in the slice.
  freq := make(map[interface{}]int)

  // this is the position in the slice starting from 0...that checks to see if
  // if it has any duplicates anywhere in the rest of the slice...
  // that is the position that iterates through the rest of the slice to be
  // compared to this before this increments to check the next element against
  // the rest of the slice.
  // NOTE: I think the inner for can be optimized to start with that := this
  //       in which case, once the first elements have been checked against the
  //       rest of the slice, when it moves to later elements, they have already
  //       been compared to the elements that came before.  
  for this := range _slice {
    for that := this; that < len(_slice); that++ {
      // If element being compared is the same...continue loop.
      // The same element is not a duplicate.
      if this == that {
        continue
      }
      // Execute compare callback.
      eq := compare(_slice[this], _slice[that])

      // If the elements were equal...
      if eq == Equal {
        // Increment the frequency _slice[this] has been seen.
        freq[_slice[this]]++
      }
      
      // If _slice[this] has been seen more than once
      if freq[_slice[this]] > 1 {
        // Flag that the slice has duplicates.
        hasDupes = true
        // Break to label won't work with the label below it.
        // So use goto to jump out of the loop.
        goto LoopDone
      }
    }
  }
  LoopDone:

  return hasDupes
}
