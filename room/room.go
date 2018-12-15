package room

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"strconv"
)

type Json struct {
	Status   int      `json:"status"`
	Result   string   `json:"result"`
	BuildNum int      `json:"build"`
	Detail   []Detail `json:"detail"`
}

type Detail struct {
	RoomName string `json:"room_name"`
	RoomNum  string `json:"room_num"`
}

func HandleRequestRoom(w http.ResponseWriter, r *http.Request) {
	var err error
	switch r.Method {
	case "GET":
		err = handleGet(w, r)
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	return
}

func handleGet(w http.ResponseWriter, r *http.Request) (err error) {

	relativePath := "jsonfiles/"
	id, err := strconv.Atoi(path.Base(r.URL.Path))

	loadJson := relativePath + strconv.Itoa(id) + ".json"

	jsonFile, err := os.Open(loadJson)
	if err != nil {
		fmt.Println("Error opening JSON file:", err)
		return
	}
	defer jsonFile.Close()

	decoder := json.NewDecoder(jsonFile)
	var jsons []Json
	for {
		err := decoder.Decode(&jsons)
		if err == io.EOF {
			break
		}
	}
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	for _, p := range jsons {

		body := Json{http.StatusOK, "ok", p.BuildNum, p.Detail,}
		returnResponse(w, body)
	}
	return

	body := Json{http.StatusNotFound, "Not Found", -1, nil}
	returnResponse(w, body)

	return
}

func returnResponse(w http.ResponseWriter, body Json) {
	res, err := json.Marshal(body)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
	return
}
