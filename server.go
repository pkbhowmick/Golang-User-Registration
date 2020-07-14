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
func main(){
	fmt.Println("Server is running...")
	http.HandleFunc("/",handler)
	http.HandleFunc("/get",get)
	http.ListenAndServe(":8080", nil)
}