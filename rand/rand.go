// nx/rand provides various sorts of random facilities for strings/bytes.
package rand

import (
  "fmt"
  "time"
  "strings"
  mrand "math/rand"
  crand "crypto/rand"
  "github.com/nu7hatch/gouuid"
)

func init() {
  // Seed math/rand on startup...it only has to be seeded once
  mrand.Seed(time.Now().UTC().UnixNano())
}

var (
  // letterRunes is the collection of characters used to select from randomly
  // when random characters are generated.
  letterRunes = []rune(
    "abcdefghijklmnopqrstuvwxyz" +
    "ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
    "0123456789" +
    "-_",
  )
)

// String generates a random string of n characters from the letterRunes
// slice above.
// It returns the generated n length string.
func String(n int) string {
  runes := make([]rune, n)
  for i := range runes {
    // I'm not sure if len(letterRunes) will ever go out of bounds here or not...
    runes[i] = letterRunes[mrand.Intn(len(letterRunes))]
  }
  return string(runes)
}

// Bytes generates a random slice of n bytes leveraging crypto/rand for the
// random bytes.
// It returns the generated n length byte slice.
func Bytes(n int) []byte {
  bytes := make([]byte, n)
  _, err := crand.Read(bytes)

  if err != nil {
    fmt.Printf("nx.rand.Bytes().error: %s\n", err.Error())
    return nil
  }

  return bytes
}

// Guid generates a UUID v4 value. But the Microsoft GUID verbage
// is used.
// It returns the generated GUID with or without hyphens depending on the value
// of the parameter removeHyphens.
func Guid(removeHyphens bool) string {
  guid, err := uuid.NewV4()

  if err != nil {
    fmt.Printf("nx.rand.Guid(%t).error: %s \n", removeHyphens, err.Error())
    return ""
  }

  var guidString string

  if removeHyphens {
    guidString = strings.Replace(guid.String(), "-", "", -1)
  } else {
    guidString = guid.String()
  }

  return guidString
}
