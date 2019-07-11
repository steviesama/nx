package rand_test

import (
  "fmt"
  "time"
  "testing"
  mrand "math/rand"
  "github.com/steviesama/nx/rand"
  "github.com/steviesama/nx/analyze"
)

func init() {
  // Seed math/rand; only needs to be seeded once.
  mrand.Seed(time.Now().UTC().UnixNano())
}

func TestString(t *testing.T) {
  genLen := 32
  genCount := 15

  // fmt.Println("Testing rand.String()")
  // fmt.Printf("Generated Length: %d\n", genLen)
  // fmt.Printf("# to Generate: %d\n", genCount)

  // Create a slice of strings.
  randStrings := make([]string, genCount)
  
  for i := range randStrings {
    randStrings[i] = rand.String(genLen)
    
    // FAIL: The following code was inserted to make sure that the test
    //       could fail. Leaving commented out in case it is needed later.
    // if i == 5 {
    //   randStrings[i] = randStrings[i-1]
    // }

    // Display the current randString.
    // fmt.Printf("randStrings[%d]: %s\n", i, randStrings[i])
  }

  // If there are duplicates in randStrings
  if analyze.HasSliceDuplicates(randStrings, nil) {
    // random strings of this length should not have duplicates
    fmt.Printf("%d rand.String(%d) were generated and duplicates were found", genCount, genLen)
    t.Errorf("%d rand.String(%d) were generated and duplicates were found", genCount, genLen)
  }
}

func TestBytes(t *testing.T) {
  genLen := 32
  genCount := 15

  // fmt.Println("Testing rand.Bytes()")
  // fmt.Printf("Generated Length: %d\n", genLen)
  // fmt.Printf("# to Generate: %d\n", genCount)

  // Create a slice of byte slices
  randBytes := make([][]byte, genCount)
  
  for i := range randBytes {
    randBytes[i] = rand.Bytes(genLen)
    
    // FAIL: The following code was inserted to make sure that the test
    //       could fail. Leaving commented out in case it is needed later.
    // if i == 5 {
    //   randBytes[i] = randBytes[i-1]
    // }

    // Display the current randBytes.
    // fmt.Printf("randBytes[%d]: %s\n", i, randBytes[i])
  }

  // If there are duplicates in randBytes
  if analyze.HasSliceDuplicates(randBytes, analyze.ByteSliceValue) {
    // random bytes of this length should not have duplicates
    fmt.Printf("%d rand.Bytes(%d) were generated and duplicates were found", genCount, genLen)
    t.Errorf("%d rand.Bytes(%d) were generated and duplicates were found", genCount, genLen)
  }
}

func TestGuid(t *testing.T) {
  // Generate a million GUIDs...tried a billion...that worked too.
  genCount := 1000000

  // fmt.Println("Testing rand.Guid()")
  // fmt.Printf("# to Generate: %d\n", genCount)

  // Create a slice of strings.
  randGuids := make([]string, genCount)
  
  for i := range randGuids {
    // Generate a GUID without hyphens
    randGuids[i] = rand.Guid(true)
    
    // FAIL: The following code was inserted to make sure that the test
    //       could fail. Leaving commented out in case it is needed later.
    // if i == 5 {
    //   randGuids[i] = randGuids[i-1]
    // }

    // Display the current randGuid.
    // fmt.Printf("randGuids[%d]: %s\n", i, randGuids[i])
  }

  // If there are duplicates in randGuids
  if analyze.HasSliceDuplicates(randGuids, nil) {
    // random strings of this length should not have duplicates
    fmt.Printf("%d rand.Guid() were generated and duplicates were found", genCount)
    t.Errorf("%d rand.Guid() were generated and duplicates were found", genCount)
  }
}