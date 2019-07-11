package analyze

import (
  "encoding/base64"
)

// ByteSliceValue takes a byte slice as an empty interface and converts it
// to []byte for processing. If a []byte isn't passed to it, it will panic.
// It returns a URL encoded string primitive representation of this.
func ByteSliceValue(this interface{}) interface{} {
  slice, ok := this.([]byte)
  if !ok {
    panic("analyze.ByteSliceValue() was not passed a []byte")
    return nil
  }
  // Return a URL encoded string in case the value is used in a URL.
  return base64.URLEncoding.EncodeToString(slice)
}
