package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
)

var UM_AP_TOKEN_URL string
var UM_AP_API_URL string
var UM_AP_GRANT_TYPE string
var UM_AP_CLIENT_ID string
var UM_AP_CLIENT_SECRET string
var UM_AP_AUDIENCE string

func main() {
    UM_AP_TOKEN_URL = os.Getenv("UM_AP_TOKEN_URL")
    UM_AP_API_URL = os.Getenv("UM_AP_API_URL")
    UM_AP_GRANT_TYPE = os.Getenv("UM_AP_GRANT_TYPE")
    UM_AP_CLIENT_ID = os.Getenv("UM_AP_CLIENT_ID")
    UM_AP_CLIENT_SECRET = os.Getenv("UM_AP_CLIENT_SECRET")
    UM_AP_AUDIENCE = os.Getenv("UM_AP_AUDIENCE")

    fmt.Println("Token Url : ",os.Getenv("UM_AP_TOKEN_URL"))
	fmt.Println("API Url : ",os.Getenv("UM_AP_API_URL"))
	fmt.Println("Grant Type : ",os.Getenv("UM_AP_GRANT_TYPE"))
	fmt.Println("Client ID : ",os.Getenv("UM_AP_CLIENT_ID"))
	fmt.Println("Client Secret : ",os.Getenv("UM_AP_CLIENT_SECRET"))
    fmt.Println("Audience : ",os.Getenv("UM_AP_AUDIENCE"))
    fmt.Println("Port : ",os.Getenv("UM_AP_PORT"))
	router := NewRouter()
    log.Fatal(http.ListenAndServe(":"+os.Getenv("UM_AP_PORT"), router))
}