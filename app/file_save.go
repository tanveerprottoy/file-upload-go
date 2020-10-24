package app

import (
	"io"
	"net/http"
	"os"
	"path/filepath"
)

// FileSave fetches the file and saves to disk
func FileSave(r *http.Request) string {
	// left shift 32 << 20 which results in 32*2^20 = 33554432
	// x << y, results in x*2^y
	err := r.ParseMultipartForm(32 << 20)
	if err != nil {
		return ""
	}
	n := r.Form.Get("name")
	// Retrieve the file from form data
	f, h, err := r.FormFile("file")
	if err != nil {
		return ""
	}
	defer f.Close()
	path := filepath.Join(".", "files")
	_ = os.MkdirAll(path, os.ModePerm)
	fullPath := path + "/" + n + filepath.Ext(h.Filename)
	file, err := os.OpenFile(fullPath, os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		return ""
	}
	defer file.Close()
	// Copy the file to the destination path
	_, err = io.Copy(file, f)
	if err != nil {
		return ""
	}
	return fullPath
}
