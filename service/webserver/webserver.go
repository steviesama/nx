// nx/service/webserver holds data types and functions related to creating a
// web server via the net/http package. It will allow the holding of various
// config data as well so the code site of the web server code can be cleaner.
package webserver

import (
  "fmt"
  "github.com/gorilla/handlers"
  "github.com/gorilla/mux"
  "net/http"
  "os"
)

// Router holds all the routes defined for the webserver that is defined
// by this package.
var Router *mux.Router

// Private slice of CORS Options.
var corsOptions []handlers.CORSOption

// Config holds all the web server config info and will be used to serialize it
// to disk via the json annotations.
type Config struct {
  // Port determines the port the web server listens on.
  Port int `json:"Port"`
  // Determines whether or not to use HTTPS.
  UseTLS bool `json:"UseTLS"`
  // The location of the cert file if UseTLS == true
  CertFile string `json:"CertFile"`
  // The location of the key file if UseTLS == true
  KeyFile string `json:"KeyFile"`
}

//--- FUNCTIONS ---//

func init() {
  // Initialize the local router.
  Router = mux.NewRouter()
}

// AddCORSOption allows the caller to specific a CORS Option to add for
// cross-origin resource sharing.
func AddCORSOption(opt handlers.CORSOption) {
  corsOptions = append(corsOptions, opt)
}

// This is a simplified version of the net/http package. If you provide only
// the info for HTTP...http.ListenAndServe() will be used...if you provide
// the SSL information, http.ListenAndServeTLS() will be used.
// It returns the return value of the ListenAndServe function used.
func ListenAndServe(config Config) error {
  // Format the port number to what http.ListenAndServe_ expects.
  portString := fmt.Sprintf(":%d", config.Port)
  corsHandler := handlers.CORS(corsOptions...)
  if config.UseTLS {
    return http.ListenAndServeTLS(
      portString,
      config.CertFile,
      config.KeyFile,
      corsHandler(handlers.LoggingHandler(os.Stdout, Router)),
    )
  } else {
    return http.ListenAndServe(
      portString,
      corsHandler(handlers.LoggingHandler(os.Stdout, Router)),
    )
  }
}
