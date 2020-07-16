package main

import (
	"log"
	"net/http"
)

func get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte (`{"message":"In get method"}`))
}

func post(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte (`{"message":"In post method"}`))
}
func put(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte (`{"message":"In put method"}`))
}
func delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte (`{"message":"In delete method"}`))
}

func main(){
	log.Println("Listening on :8080...")
	
	index := http.FileServer(http.Dir("./static"))
    http.Handle("/", index)

	http.HandleFunc("/api/get",get)
	http.HandleFunc("/api/post",post)
	http.HandleFunc("/api/put",put)
	http.HandleFunc("/api/delete",delete)
	http.ListenAndServe(":8080", nil)
}