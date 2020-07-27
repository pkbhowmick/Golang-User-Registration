package main

import (
	"context"
	"time"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"encoding/json"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
	"github.com/dgrijalva/jwt-go"
)


var SECRET_KEY = []byte("gosecretkey")

type User struct{
	FirstName string `json:"firstname" bson:"firstname"`
	LastName string `json:"lastname" bson:"lastname"`
	Email string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
}

var client *mongo.Client

func getHash(pwd []byte) string {
    
    hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
    if err != nil {
        log.Println(err)
    }
    return string(hash)
}

func GenerateJWT()(string,error){
	token:= jwt.New(jwt.SigningMethodHS256)

	tokenString, err :=  token.SignedString(SECRET_KEY)

	if err !=nil{
		log.Println("Error in JWT token generation")
		return "",err
	}

	return tokenString, nil
}


func getUsers(response http.ResponseWriter, request *http.Request){
  response.Header().Set("Content-Type","application/json")
  var users []User
  collection:= client.Database("GODB").Collection("user")
  ctx,_ := context.WithTimeout(context.Background(),10*time.Second)
  cursor,err:=collection.Find(ctx,bson.M{})

  if err!=nil{
	  response.WriteHeader(http.StatusInternalServerError)
	  response.Write([]byte(`{"message":"`+err.Error()+`"}`))
	  return
  }
  defer cursor.Close(ctx)
  for cursor.Next(ctx){
	  var user User
	  cursor.Decode(&user)
	  users = append(users, user)
  }

  if err:=cursor.Err(); err!=nil{
	response.WriteHeader(http.StatusInternalServerError)
	response.Write([]byte(`{"message":"`+err.Error()+`"}`))
	return
  }
  json.NewEncoder(response).Encode(users)

}

func userLogin(response http.ResponseWriter, request *http.Request){
  response.Header().Set("Content-Type","application/json")
  var user User
  var dbuser User
  json.NewDecoder(request.Body).Decode(&user)
  collection:= client.Database("GODB").Collection("user")
  ctx,_ := context.WithTimeout(context.Background(),10*time.Second)
  err:= collection.FindOne(ctx, bson.M{"email":user.Email}).Decode(&dbuser)

  if err!=nil{
	  response.WriteHeader(http.StatusInternalServerError)
	  response.Write([]byte(`{"message":"`+err.Error()+`"}`))
	  return
  }
  userPass:= []byte(user.Password)
  dbPass:= []byte(dbuser.Password)

  passErr:= bcrypt.CompareHashAndPassword(dbPass, userPass)

  if passErr != nil{
	  log.Println(passErr)
	  response.Write([]byte(`{"response":"Wrong Password!"}`))
	  return
  }
  jwtToken, err := GenerateJWT()
  if err != nil{
	response.WriteHeader(http.StatusInternalServerError)
	response.Write([]byte(`{"message":"`+err.Error()+`"}`))
	return
  }
  response.Write([]byte(`{"token":"`+jwtToken+`"}`))
  
}

func userSignup(response http.ResponseWriter, request *http.Request){
	response.Header().Set("Content-Type","application/json")
	var user User
	json.NewDecoder(request.Body).Decode(&user)
	user.Password = getHash([]byte(user.Password))
	collection := client.Database("GODB").Collection("user")
	ctx,_ := context.WithTimeout(context.Background(), 10*time.Second)
	result,_ := collection.InsertOne(ctx,user)
	json.NewEncoder(response).Encode(result)
}

func updateUser(response http.ResponseWriter, request *http.Request){
	response.Header().Set("Content-Type","application/json")
}

func deleteUser(response http.ResponseWriter, request *http.Request){
	response.Header().Set("Content-Type","application/json")
}



func main(){
	log.Println("Starting the application")

	router:= mux.NewRouter()
	ctx,_ := context.WithTimeout(context.Background(), 10*time.Second)
	client,_= mongo.Connect(ctx,options.Client().ApplyURI("mongodb://localhost:27017"))

	router.HandleFunc("/api/users",getUsers).Methods("GET")
	router.HandleFunc("/api/user/login",userLogin).Methods("POST")
	router.HandleFunc("/api/user/signup",userSignup).Methods("POST")
	router.HandleFunc("/api/user/{email}",updateUser).Methods("PUT")
	router.HandleFunc("/api/user/{email}",deleteUser).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))

}
