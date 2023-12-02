package services

import (
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

type FileStorage interface {
	Save(file multipart.File, fileName, filePath string) (int, error)
	Delete(fileName, filePath string) (int, error)
	GetRelative(fileName, filePath string) (string, string, int, error)
	GetAbsolute(fileName, filePath string) (string, int, error)
}

type LocalFileStorage struct {
	Path string
}

func NewLocalFileStorage(filePath string) *LocalFileStorage {
	return &LocalFileStorage{
		Path: filePath,
	}
}

func (data *LocalFileStorage) GetRelative(fileName, filePath string) (string, string, int, error) {
	if filePath == "" {
		filePath = "temp"
	}
	filePathCombined := fmt.Sprintf("%s/%s", filePath, fileName)
	path := filepath.Join(data.Path, filePathCombined)

	// Check if the file exists
	if _, err := os.Stat(path); err != nil {
		return "", "", http.StatusNotFound, err
	}

	return fileName, filePath, 0, nil
}

func (data *LocalFileStorage) GetAbsolute(fileName, filePath string) (string, int, error) {
	if filePath == "" {
		filePath = "temp"
	}
	savePath := fmt.Sprintf("%s/%s/%s", data.Path, filePath, fileName)

	absPath, _ := filepath.Abs(savePath)

	// Check if the file exists
	if _, err := os.Stat(absPath); err != nil {
		return "", http.StatusNotFound, err
	}

	return absPath, 0, nil
}

func (data *LocalFileStorage) Save(file multipart.File, fileName, filePath string) (int, error) {

	if filePath == "" {
		filePath = "temp"
	}
	savePath := fmt.Sprintf("%s/%s", data.Path, filePath)

	absPath, err := filepath.Abs(savePath)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	if _, err := os.Stat(absPath); os.IsNotExist(err) {
		if err := os.MkdirAll(absPath, 0777); err != nil {
			return http.StatusInternalServerError, err
		}
	}

	saveFilePath := fmt.Sprintf("%s/%s", absPath, fileName)

	dst, err := os.Create(saveFilePath)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	defer dst.Close()

	// Copy the contents of the uploaded file to the destination file
	_, err = io.Copy(dst, file)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return 0, nil
}

func (data *LocalFileStorage) Delete(fileName, filePath string) (int, error) {
	if filePath == "" {
		filePath = "temp"
	}
	fullFilePath := fmt.Sprintf("%s/%s/%s", data.Path, filePath, fileName)

	// Check if the file exists
	if _, err := os.Stat(fullFilePath); err == nil {
		// File exists, so delete it
		err := os.Remove(fullFilePath)
		if err != nil {
			fmt.Println("Error deleting the file:", err)
			return http.StatusInternalServerError, err
		}

		fmt.Printf("File %s deleted successfully\n", fullFilePath)
	} else if os.IsNotExist(err) {
		return http.StatusNotFound, fmt.Errorf("file %s does not exist", fullFilePath)
	} else {
		return http.StatusInternalServerError, fmt.Errorf("error checking if the file exists: %s", err.Error())
	}

	return 0, nil
}
