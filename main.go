package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	//"path"
	//"strconv"
)

type Json struct {
	Status int      `json:"status"`
	Result string   `json:"result"`
	Id     int      `json:"id"`
	//Nya    []string `json:"nya"`
}

type Building struct {
	Id int `json:"id"`
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
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

func handleRequest(w http.ResponseWriter, r *http.Request) {
	var err error
	switch r.Method {
	case "GET":
		err = handleGet(w, r)
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func handleGet(w http.ResponseWriter, r *http.Request) (err error) {
	//id, err := strconv.Atoi(path.Base(r.URL.Path))

	//fmt.Println(id)
	//body := Json{http.StatusOK, "ok", id}
	jsonFile, err := os.Open("building.json")
	if err != nil {
		fmt.Println("Error opening JSON file:", err)
		return
	}
	defer jsonFile.Close()

	decoder := json.NewDecoder(jsonFile)
	var building []Building
	for {
		err := decoder.Decode(&building)
		if err == io.EOF {
			break
		}
	}
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}
	for _, p := range building {
		fmt.Printf("%d : %f\n", p.Id, p.Lat)
	}
	fmt.Println(building)
	//returnResponse(w, body)

	return
}

//func buildingJsonHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
//
//}

func main() {
	http.HandleFunc("/api/building/", handleRequest)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe", err)
	}
}
