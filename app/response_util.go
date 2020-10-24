package app

import (
	"encoding/json"
	"net/http"
)

func writeResponse(w http.ResponseWriter, bytes []byte) {
	_, _ = w.Write(bytes)
}

func Respond(w http.ResponseWriter, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		RespondError(w, err, http.StatusInternalServerError)
		return
	}
	writeResponse(w, response)
}

func RespondError(w http.ResponseWriter, err error, code int) {
	response, err := json.Marshal(map[string]string{"error": err.Error()})
	w.WriteHeader(code)
	if err != nil {
		writeResponse(w, []byte(err.Error()))
		return
	}
	writeResponse(w, response)
}
