package schemes

import "net/http"

func HandleFind(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("find scheme"))
}
