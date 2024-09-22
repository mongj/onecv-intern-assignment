package applicants

import "net/http"

func HandleCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("create applicant"))
}
