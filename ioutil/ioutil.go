// nx/ioutil provides an abbreviated way to handle reading & writing.
package ioutil

import (
  "io/ioutil"
)

// FwriteFile takes a filepath and a contents byte slice and writes the contents
// into a file at filepath.
// It returns a boolean indicating whether the operation was successful.
func FwriteFile(filepath string, contents []byte) bool {
  err := ioutil.WriteFile(filepath, contents, 0666)
  if err != nil {
    return false
  }
  return true
}

// FreadFile takes a filepath and reads the contents of a text file.
// It returns either the string contents of the file...or the string
// message for the error if one occurred.
func FreadFile(filepath string) string {
  textBytes, err := ioutil.ReadFile(filepath)
  if err != nil {
    return err.Error()
  }
  return string(textBytes)
}