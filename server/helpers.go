package server

import (
	"encoding/json"
	"log"
	"net/http"
)

type errorMsg struct {
	Message string `json:"details,omitempty"`
}

func writeJSON(w http.ResponseWriter, code int, d interface{}) {
	switch v := d.(type) {
	case error:
		if code >= 500 {
			log.Printf("%v", v)
		}
		d = errorMsg{Message: v.Error()}
	}

	byt, err := json.Marshal(d)
	if err != nil {
		log.Printf("json.Marshal: %v", err)
	}
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(code)
	if _, err := w.Write(byt); err != nil {
		log.Printf("w.Write: %v", err)
	}
}
