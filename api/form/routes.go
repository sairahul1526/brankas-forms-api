package form

import (
	CONSTANT "forms-api/constant"
	"os"

	"github.com/gorilla/mux"
)

// LoadFormRoutes - load all form routes with form prefix
func LoadFormRoutes(router *mux.Router) {
	formRoutes := router.PathPrefix("/form").Subrouter()

	// create image directory
	os.Mkdir(CONSTANT.ImageFolderPath, os.ModePerm)

	// upload
	formRoutes.HandleFunc("", FormGet).Methods("GET")
	formRoutes.HandleFunc("/upload", FormUpload).Methods("POST")

}
