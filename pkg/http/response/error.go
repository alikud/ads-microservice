package response

import (
	"encoding/json"
	"net/http"
)

type errorResponse struct {
	Message string `json:"message"`
}

func Error(w http.ResponseWriter, message string, status int) {
	res := errorResponse{
		Message: message,
	}

	b, err := json.Marshal(res)
	if err != nil {
		Error(w, "Something went wrong", http.StatusInternalServerError)
	}

	w.WriteHeader(status)
	w.Header().Add("Content-Type", "application/json")
	_, _ = w.Write(b)
}
