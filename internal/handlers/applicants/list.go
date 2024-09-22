package applicants

import "net/http"

func HandleList(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("list applicants"))
}
