package app

import (
	"errors"
	"net/http"
)

// FileHandler handles file specific requests
type FileHandler struct {
}

// PostFile posts new file
func (f *FileHandler) PostFile(w http.ResponseWriter, r *http.Request) {
	path := FileSave(r)
	if path == "" {
		RespondError(w, errors.New("an error occurred"), http.StatusInternalServerError)
		return
	}
	Respond(w, map[string]string{"path": path})
}
