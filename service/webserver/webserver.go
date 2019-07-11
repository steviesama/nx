// nx/webserver holds data types and functions related to created a web server
// via the net/http package. It will allow the holding of various config data
// as well so the code site of the web server code can be cleaner.
package webserver

import (
  "fmt"
  "net/http"
)

// Config holds all the web server config info and will be used to serialize it
// to disk via the json annotations.
type Config struct {
  // Port determines the port the web server listens on.
  Port int
  // Determines whether or not to use HTTPS.
  UseTLS bool `json:"UseTLS"`
  // The location of the cert file is UseTLS == true
  CertFile string `json:"CertFile"`
  // The location of the key file is UseTLS == true
  KeyFile string `json:"KeyFile"`
}

// This is a simplified version of the net/http package. If you provide only
// the info for HTTP...http.ListenAndServe() will be used...if you provide
// the SSL information, http.ListenAndServeTLS() will be used.
// It returns the return value of the ListenAndServe function used.
func ListenAndServe(config Config) error {
  portString := fmt.Sprintf(":%d", config.Port)
  if config.UseTLS {
    return http.ListenAndServeTLS(
      portString,
      config.CertFile,
      config.KeyFile,
      nil, //replace this with handlers once added.
    )
  } else {
    return http.ListenAndServe(
      portString,
      nil, //replace this with handlers once added.
    )
  }
}
