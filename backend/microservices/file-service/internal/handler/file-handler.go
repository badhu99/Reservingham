package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"path"

	"github.com/badhu99/file-service/internal/dto"
	"github.com/badhu99/file-service/internal/services"
	"github.com/badhu99/file-service/internal/utility"
	"github.com/gorilla/mux"
)

type HandlerData struct {
	FileStorage services.FileStorage
}

func (data *HandlerData) Upload(w http.ResponseWriter, r *http.Request) {

	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		return
	}

	file, handler, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Error retrieving the file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	fileName := r.FormValue("file_name")
	if fileName == "" {
		fileName = handler.Filename
	} else {
		fileName = fmt.Sprintf("%s%s", fileName, path.Ext(handler.Filename))
	}

	filePath := r.FormValue("file_path")

	statusCode, err := data.FileStorage.Save(file, fileName, filePath)
	if err != nil {
		http.Error(w, err.Error(), statusCode)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (data *HandlerData) DeleteFile(w http.ResponseWriter, r *http.Request) {
	inputData := dto.FileData{}

	err, statusCode := utility.ValidateBody(&inputData, r.Body)
	if err != nil {
		http.Error(w, err.Error(), statusCode)
		return
	}

	statusCode, err = data.FileStorage.Delete(inputData.FileName, inputData.FilePath)
	if err != nil {
		http.Error(w, err.Error(), statusCode)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (data *HandlerData) GetFile(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fileName := vars["fileName"]

	filePath := r.URL.Query().Get("filePath")

	fileName, filePath, statusCode, err := data.FileStorage.GetRelative(fileName, filePath)
	if err != nil {
		http.Error(w, err.Error(), statusCode)
		return
	}
	url := fmt.Sprintf("http://localhost:8084/file/access/%s?filePath=%s", fileName, filePath)
	returnData := dto.FileData{
		FilePath: url,
	}

	jsonReturnData, _ := json.Marshal(returnData)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonReturnData)
}

func (data *HandlerData) ServeFile(w http.ResponseWriter, r *http.Request) {

	companyId, ok := r.Context().Value("companyId").(string)
	if !ok {
		http.Error(w, "CompanyId not found", http.StatusBadRequest)
		return
	}

	userId, ok := r.Context().Value("userId").(string)
	if !ok {
		http.Error(w, "UserId not found", http.StatusBadRequest)
		return
	}

	log.Printf("CompanyId: %s, UserId: %s \n", companyId, userId)
	vars := mux.Vars(r)
	fileName := vars["fileName"]

	filePath := r.URL.Query().Get("filePath")

	// Construct the file path using the specified upload directory
	fullPath, statusCode, err := data.FileStorage.GetAbsolute(fileName, filePath)
	if err != nil {
		http.Error(w, err.Error(), statusCode)
		return
	}

	// Serve the image file
	http.ServeFile(w, r, fullPath)
}
