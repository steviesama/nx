// nx/ioutil provides an abbreviated way to handle reading & writing.
package ioutil

import (
  "io/ioutil"
)

// WriteFile takes a filepath and a contents byte slice and writes the contents
// into a file at filepath.
// It returns a boolean indicating whether the operation was successful.
func WriteFile(filepath string, contents []byte) bool {
  err := ioutil.WriteFile(filepath, contents, 0666)
  if err != nil {
    return false
  }
  return true
}

// ReadFile takes a filepath and reads the contents of a text file.
// It returns either the string contents of the file...or the string
// message for the error if one occurred.
func ReadFile(filepath string) string {
  textBytes, err := ioutil.ReadFile(filepath)
  if err != nil {
    return err.Error()
  }
  return string(textBytes)
}

// ReadFileBytes takes a file path and reads the contents of the file specified
// by it. Panics if the file doesn't exist.
// It returns the contents of the file read as a slice of bytes.
// TODO: write the code
func ReadFileBytes(filepath string) []byte {
  return nil
}
