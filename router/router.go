package router

import (
	"net/http"
	"os"
	"path/filepath"
	"github.com/gorilla/mux"
	"github.com/nouveau05/nouveau-microservices-go/middleware"
)

// spaHandler implements the http.Handler interface, so we can use it
// to respond to HTTP requests. The path to the static directory and
// path to the index file within that static directory are used to
// serve the SPA in the given static directory.
type spaHandler struct {
	staticPath string
	indexPath  string
}

// ServeHTTP inspects the URL path to locate a file within the static dir
// on the SPA handler. If a file is found, it will be served. If not, the
// file located at the index path on the SPA handler will be served. This
// is suitable behavior for serving an SPA (single page application).
func (h spaHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    // get the absolute path to prevent directory traversal
	path, err := filepath.Abs(r.URL.Path)
	if err != nil {
        // if we failed to get the absolute path respond with a 400 bad request
        // and stop
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

    // prepend the path with the path to the static directory
	path = filepath.Join(h.staticPath, path)

    // check whether a file exists at the given path
	_, err = os.Stat(path)
	if os.IsNotExist(err) {
		// file does not exist, serve index.html
		http.ServeFile(w, r, filepath.Join(h.staticPath, h.indexPath))
		return
	} else if err != nil {
        // if we got an error (that wasn't that the file doesn't exist) stating the
        // file, return a 500 internal server error and stop
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

    // otherwise, use http.FileServer to serve the static dir
	http.FileServer(http.Dir(h.staticPath)).ServeHTTP(w, r)
}

// Router is exported and used in main.go
func Router() *mux.Router {

	router := mux.NewRouter()

	// Route information
	router.HandleFunc("/api/venture/{id}", middleware.GetVenture).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/venture", middleware.GetAllVentures).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/newventure", middleware.CreateVenture).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/venture/{id}", middleware.UpdateVenture).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/venture/{id}", middleware.DeleteVenture).Methods("DELETE", "OPTIONS")

	// static files
	spa := spaHandler{staticPath: "app/ui", indexPath: "index.html"}
	router.PathPrefix("/").Handler(spa)

	return router
}
