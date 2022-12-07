package server

import "net/http"

func (a *API) ping(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "405 Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Write([]byte("pong"))
}
