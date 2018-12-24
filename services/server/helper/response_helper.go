package helper

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/lib/pq"
	log "github.com/sirupsen/logrus"
)

const (
	ContentJson = "application/json"
	ContentText = "text/plain"
)

func ResponseHelper(w http.ResponseWriter, http_status int, content_type string, body string) {
	w.Header().Set("Content-Type", content_type)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	w.WriteHeader(http_status)
	fmt.Println("HERE")
	w.Write([]byte(body))
}

func ResponseError(caller string, w http.ResponseWriter, err error) {
	if err == nil {
		return
	}
	log.Error(err.Error())
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	if httpErr, ok := err.(*HttpError); ok && httpErr.Status != http.StatusInternalServerError {
		msg, _ := json.Marshal(httpErr)
		http.Error(w, string(msg), httpErr.Status)
	} else {
		if pqErr, ok := err.(*pq.Error); ok {
			log.Println(pqErr.Code)
			log.Println(pqErr.Message)
		}
		http.Error(w, `{"status": 500, "message": "Internal Server Error"}`, http.StatusInternalServerError)
	}

}
