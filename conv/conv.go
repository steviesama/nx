// nx/conv provides various kinds of conversion that the Go built-in library does not.
package conv

import (
  "reflect"
)

// InterfaceSlice takes a slice of any kind stored as an empty interface
// and transforms it into an empty interface slice.
// It returns the empty interface slice after transformation unless the
// caller didn't pass a slice argument...in which case it will panic.
func InterfaceSlice(slice interface{}) []interface{} {
  s := reflect.ValueOf(slice)
  if s.Kind() != reflect.Slice {
    // TODO: change this so that it uses errors.New
    panic("nx.conv.InterfaceSlice().error: given a non-slice type")
  }

  newSlice := make([]interface{}, s.Len())

  for i := 0; i < s.Len(); i++ {
    newSlice[i] = s.Index(i).Interface()
  }

  return newSlice
}
