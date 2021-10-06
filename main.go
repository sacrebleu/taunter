/*
	taunts implements scorn-as-a-service, a lightweight REST API that generates scornful responses for use in slack or
    other chat clients.
*/
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"runtime"
	"sacrebleu/saas/taunts"
)

var (
	_, b, _, _ = runtime.Caller(0)
	basepath   = filepath.Dir(b)
	datapath   = taunts.Paths{
		Prefixes: fmt.Sprintf("%s/data/prefix.txt", basepath),
		Taunts: fmt.Sprintf("%s/data/taunts.txt", basepath),
		Verbs: fmt.Sprintf("%s/data/verbs.txt", basepath),
	}
	data 	   = taunts.LoadData(datapath)
)


func index(writer http.ResponseWriter, request *http.Request){
	keys, ok := request.URL.Query()["target"]

	if !ok || len(keys[0]) < 1 {
		log.Println("Url Param 'target' is missing")
		return
	}

	var taunt = taunts.Generate(keys[0], data)
	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Content-Type", "application/json")

	log.Printf("[INFO] Taunting %s", keys[0])
	json.NewEncoder(writer).Encode(taunt)
	return
}

func handleRequests() {
	http.HandleFunc("/", index)
	log.Fatal(http.ListenAndServe(":8080", nil))
}


func main() {
	handleRequests()
}
