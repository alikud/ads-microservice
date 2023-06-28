package response

import (
	"encoding/json"
	"net/http"
)

func Success(w http.ResponseWriter, data interface{}, status int) error {
	b, err := json.Marshal(data)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, err = w.Write(b)
	if err != nil {
		return err
	}

	return nil
}
