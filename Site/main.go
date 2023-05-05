package main

import (
	"net/http"
	"html/template"
	"io/ioutil"
	"encoding/json"
	"strconv"
	"fmt"
)

type BlockchainData struct {
    Id int `json:"id"`
    Data Blockchain `json:"blockchain"`
}

type Block struct {
    PreviousBlockHash []byte
    CurrentBlockHash []byte
    UserId int
    OtherId int
    Time int64
    OtherTime int64
}

type Blockchain struct {
    Blocks []*Block
}

func main() {
	fs := http.FileServer(http.Dir("./"))
	http.Handle("/", http.StripPrefix("/", fs))
	http.HandleFunc("/home", Handler)
	http.ListenAndServe(":8080", nil)
	fmt.Println("Starting localhost:8080")
}

func JsonToBlockchains(data []byte) []BlockchainData{
    blockchains := []BlockchainData{}
    json.Unmarshal(data, &blockchains)

    return blockchains
}

func GetBlockchainByID(id int64) BlockchainData {
	data, _ := ioutil.ReadFile("./blockchains.json")
	dataStruct := JsonToBlockchains(data)

	return dataStruct[id]
}

func Handler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./index.html"))

	switch r.Method {
	case "GET":
		fmt.Println("GET")
	case "POST": // Gestion d'erreur
		if err := r.ParseForm(); err != nil {
			return
		}
	}

	id, err := strconv.ParseInt(r.URL.Query().Get("id"), 10, 64)
	fmt.Println(id)
	if err != nil {
		fmt.Println(err)
	}

	data := GetBlockchainByID(id)

	err = tmpl.Execute(w, data)

	if err != nil {
		fmt.Println(err)
	}
}