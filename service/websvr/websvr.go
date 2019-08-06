// nx/service/websvr holds data types and functions related to creating a
// web server via the net/http package. It will allow the holding of various
// config data as well so the code site of the web server code can be cleaner.
// TODO: Need to fine-tune this package so it has more flexibility on CORSOptions
// 			 and more abstraction for routing.
package websvr

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// Router holds all the routes defined for the websvr that is defined
// by this package. The caller should access this via websvr.Router and
// add routes manually outside this package.
var Router *mux.Router

// Private slice of CORS Options.
var corsOptions []handlers.CORSOption

var (
	// Preset to allow cross-origin request scripting from all origins.
	AllowAllOriginsCORSOption handlers.CORSOption
	// Preset to allow all HTTP methods to be accepted by all handlers.
	AllowAllMethodsCORSOption handlers.CORSOption
)

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
	// Setup CORSOption presets.
	AllowAllOriginsCORSOption = handlers.AllowedOrigins([]string{"*"})
	AllowAllMethodsCORSOption = handlers.AllowedMethods([]string{"HEAD", "GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"})
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
	// I don't think this defer can execute...seems like when I CTRL+C'd to kill
	// the program...I didn't see it print to the console.
	// TODO: Remove defer or process the kill signal so this defer can print.
	defer fmt.Printf("\n...closing web server...\n")
	fmt.Printf("Listening on port %d...", config.Port)
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
