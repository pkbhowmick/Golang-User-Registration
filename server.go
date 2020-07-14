package main

import (
	"fmt"
	"net/http"
)

func handler(writer http.ResponseWriter, request *http.Request){
	fmt.Fprintf(writer, "Hello World")

}

func get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte (`{"message":"In get method"}`))
}

func post(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte (`{"message":"In post method"}`))
}
func put(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte (`{"message":"In put method"}`))
}
func delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte (`{"message":"In delete method"}`))
}

func main(){
	fmt.Println("Server is running...")
	http.HandleFunc("/",handler)
	http.HandleFunc("/api/get",get)
	http.HandleFunc("/api/post",post)
	http.HandleFunc("/api/put",put)
	http.HandleFunc("/api/delete",delete)
	http.ListenAndServe(":8080", nil)
}