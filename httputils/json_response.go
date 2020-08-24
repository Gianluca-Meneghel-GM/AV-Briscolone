package httputils

import (
	"encoding/json"
	"log"
	"net/http"
)

func WriteJsonResponse(writer http.ResponseWriter, object interface{}) {
	enc := json.NewEncoder(writer)
	writer.Header().Set("Content-Type", "application/json")
	err := enc.Encode(object)
	if err != nil {
		log.Println(err)
		http.Error(writer, "could not serialize object", http.StatusInternalServerError)
	}
}
