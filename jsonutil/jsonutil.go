// nx/jsonutil provides some json utilities for marshalling and unmarshalling
// without the caller having to deal with the errors directly. Additionally,
// it offers a means to save/load datatypes to and from files without having
// to reinvent the wheel.
// 
// Be sure to add json annotations to all fields that need to be saved/loaded
// so the encoding/json package and reflect on the data type correctly.
// 
// Example:
// type LoginInfo struct {
//   Username string `json:"user_login"`
//   Password string `json:"user_pass"`
// }
// 
// The json annotation names do not have to match the name of the Go struct fields.
// As you can see in the example...the names reflect something you might find in
// a SQL Database definition.
package jsonutil

import (
  "fmt"
  "reflect"
  "encoding/json"
  "github.com/steviesama/nx/ioutil"
)

// Unmarshal takes a data string representing the JSON object to decode into
// the empty interface v. It returns a boolean indicating whether or not the
// process was successful.
func Unmarshal(data string, v interface{}) bool {
  jsonErr := json.Unmarshal([]byte(data), &v)

  // if there was an error, log it
  if jsonErr != nil {
    fmt.Printf("====jsonErr: %s\n", jsonErr.Error())
    return false
  }

  return true
}

// MarshalBytes takes an empty interface v holding the data type which should be
// marshalled to a byte slice and returned and a boolean isArray indicating
// whether or not v is an array so the appropriate empty object or array string
// can be set in the event of an error.
// It returns the marshalled data as a slice of bytes.
func MarshalBytes(v interface{}) []byte {
  var useOnError string

  // Use reflection to determine whether or not v is a slice or array
  switch reflect.TypeOf(v).Kind() {
  case reflect.Array:
    fallthrough
  case reflect.Slice:
    // Similar to the case above...except isArray gives the caller the
    // change to specific whether or not it was an array so the expected
    // type is embedded in the json data.
    useOnError = "[]"
  default:
    // When there is an error marshalling...use a singular empty object
    // to return to the user. It can cause some issues otherwise
    useOnError = "{}"
  }

  jsonBytes, err := json.Marshal(v)

  // If there is an error...apply the useOnError string
  if err != nil {
    jsonBytes = []byte(useOnError)
  }

  return jsonBytes
}

// Marshal simply converts the []byte returned from MarshalBytes to a string
// and returns it to the caller. Refer to the MarshalBytes comments for details.
// It returns the marshalled data in string form.
func Marshal(v interface{}) string {
  return string(MarshalBytes(v))
}

// MarshalIndentBytes work exactly like MarshalBytes except it accepts a prefix
// & indent string which it applies to the output so it is readable and not on
// a single line.
// It returns the marshalled data as a slice of bytes.
func MarshalIndentBytes(v interface{}, prefix, indent string) []byte {
  var useOnError string

  // Use reflection to determine whether or not v is a slice or array
  switch reflect.TypeOf(v).Kind() {
  case reflect.Array:
    fallthrough
  case reflect.Slice:
    // Similar to the case above...except isArray gives the caller the
    // change to specific whether or not it was an array so the expected
    // type is embedded in the json data.
    useOnError = "[]"
  default:
    // When there is an error marshalling...use a singular empty object
    // to return to the user. It can cause some issues otherwise
    useOnError = "{}"
  }

  jsonBytes, err := json.MarshalIndent(v, prefix, indent)

  // If there is an error...apply the useOnError string
  if err != nil {
    jsonBytes = []byte(useOnError)
  }

  return jsonBytes
}

// MarshalIndent works exactly like Marshal except it accepts a prefix & indent
// string which it applies to the output so it is readable and not on a single line.
// It returns the marshalled data in string form.
func MarshalIndent(v interface{}, prefix, indent string) string {
  return string(MarshalIndentBytes(v, prefix, indent))
}

// LoadFromFile takes a filepath and an empty interface v. It reads the data from
// filepath into the datatype in v.
func LoadFromFile(filepath string, v interface{}) {
  Unmarshal(ioutil.FreadFile(filepath), &v)
}

// SaveToFile takes a filepath and a datatype wrapped in an empty interface which
// it then uses MarshalIndentBytes to write it's data to the filepath with the util.
// It returns the return value of ioutil.FwriteFile to determine success.
func SaveToFile(filepath string, v interface{}) bool {
  return ioutil.FwriteFile(filepath, MarshalIndentBytes(v, "", "  "))
}
