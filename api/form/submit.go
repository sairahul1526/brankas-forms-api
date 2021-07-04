package form

import (
	"fmt"
	CONFIG "forms-api/config"
	CONSTANT "forms-api/constant"
	DB "forms-api/database"
	"html/template"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"

	UTIL "forms-api/util"
)

// FormGet godoc
// @Tags Form
// @Summary Get form HTML
// @Router /form [get]
// @Produce text/html
// @Success 200
func FormGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	var response = make(map[string]interface{})

	// parse html template
	tpl, err := template.ParseFiles(CONSTANT.FormUploadHTMLFile)
	if err != nil {
		fmt.Println("FormGet", err)
		UTIL.SetReponse(w, CONSTANT.StatusCodeServerError, err.Error(), "", CONSTANT.ShowDialog, response)
		return
	}

	// replace variables in template
	err = tpl.Execute(w, map[string]string{
		"Auth": CONFIG.AuthFormKey,
	})
	if err != nil {
		fmt.Println("FormGet", err)
		UTIL.SetReponse(w, CONSTANT.StatusCodeServerError, err.Error(), "", CONSTANT.ShowDialog, response)
		return
	}
}

// FormUpload godoc
// @Tags Form
// @Summary Upload form file
// @Router /form/upload [post]
// @Param file formData file true "File to be uploaded"
// @Param auth formData string true "Auth key from server"
// @Accept multipart/form-data
// @Produce json
// @Success 200
func FormUpload(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var response = make(map[string]interface{})

	// set max upload size
	r.Body = http.MaxBytesReader(w, r.Body, CONSTANT.MaxFileUploadSize)

	// set max buffer size
	err := r.ParseMultipartForm(CONSTANT.MaxFileUploadSize)
	if err != nil {
		fmt.Println("FormUpload", err)
		UTIL.SetReponse(w, CONSTANT.StatusCodeForbidden, err.Error(), "", CONSTANT.ShowDialog, response)
		return
	}

	// get file handler
	file, handler, err := r.FormFile("file")
	if err != nil {
		fmt.Println("FormUpload", err)
		UTIL.SetReponse(w, CONSTANT.StatusCodeForbidden, err.Error(), "", CONSTANT.ShowDialog, response)
		return
	}

	// check if file is image
	if !strings.HasPrefix(UTIL.GetFileMIMEType(filepath.Ext(handler.Filename)), "image/") {
		UTIL.SetReponse(w, CONSTANT.StatusCodeForbidden, "", "Upload only image", CONSTANT.ShowDialog, response)
		return
	}

	// check if auth key is valid
	if !strings.EqualFold(r.FormValue("auth"), CONFIG.AuthFormKey) {
		UTIL.SetReponse(w, CONSTANT.StatusCodeForbidden, "", "Auth key is not correct", CONSTANT.ShowDialog, response)
		return
	}

	// parse and save file
	if file != nil {
		defer file.Close()

		fileName, fileSize, err := UTIL.SaveToDisk(CONSTANT.ImageFolderPath, file, filepath.Ext(handler.Filename))
		if err != nil {
			fmt.Println("UploadFile", err)
			UTIL.SetReponse(w, CONSTANT.StatusCodeForbidden, err.Error(), "", CONSTANT.ShowDialog, response)
			return
		}

		// add file meta to database
		DB.InsertSQL(CONSTANT.FormSubmissionsTable, map[string]string{
			"name":        fileName,
			"mime_type":   UTIL.GetFileMIMEType(filepath.Ext(handler.Filename)),
			"size":        strconv.FormatInt(fileSize, 10),
			"uploaded_at": UTIL.GetCurrentTime().Format("2006-01-02 15:04:05"),
		})

		response["file"] = fileName
	}

	UTIL.SetReponse(w, CONSTANT.StatusCodeOk, "", "", CONSTANT.ShowDialog, response)
}
