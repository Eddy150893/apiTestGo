package response
import (
	"encoding/json"
	"net/http"
)
func Response(w http.ResponseWriter, status int, results string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(results)
}
