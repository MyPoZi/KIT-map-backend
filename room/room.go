package room

import "net/http"


func HandleRequestRoom(w http.ResponseWriter, r *http.Request) {
	var err error
	switch r.Method {
	case "GET":
		//err = handleGet(w, r)
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	return
}

