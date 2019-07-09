package rand_test

import (
  "fmt"
  "time"
  "testing"
  mrand "math/rand"
  "github.com/steviesama/nx/analyze"
  "github.com/steviesama/nx/rand"
)

func init() {  
  mrand.Seed(time.Now().UTC().UnixNano())
}

func TestString(t *testing.T) {
  genLen := 32
  genCount := 15

  fmt.Println("Testing rand.String()")
  fmt.Printf("Generated Length: %d\n", genLen)
  fmt.Printf("# to Generate: %d\n", genCount)

  randStrings := make([]string, genCount)
  
  for i := range randStrings {
    randStrings[i] = rand.String(genLen)
    // Display the current randString.
    fmt.Printf("randStrings[%d]: %s\n", i, randStrings[i])
  }

  // If there are duplicates in randStrings
  if analyze.HasSliceDuplicates(randStrings, analyze.StringCompare) {
    // random strings of this length should not have duplicates
    t.Errorf("%d rand.String(%d) were generated and duplicates were found", genCount, genLen)
  }
}

func TestBytes(t *testing.T) {

}

// func TestGuid(t *testing.T) {

// }