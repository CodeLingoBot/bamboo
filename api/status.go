package api

import (
	"io"
	"net/http"
)

// HandleStatus: Status Handler
func HandleStatus(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "OK")
}
