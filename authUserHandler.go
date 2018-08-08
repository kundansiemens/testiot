package main

import (
    "fmt"
    "net/http"
    "github.com/gorilla/mux"
    "io/ioutil"
    "bytes"
    "encoding/json"
    "time"
)

var access_token string
var token_type string
var expires_in int64

func generateToken() string{
    fmt.Println("getToken called")
    
    var body = []byte(`{
        "grant_type": "`+UM_AP_GRANT_TYPE+`",
        "client_id": "`+UM_AP_CLIENT_ID+`",
        "client_secret": "`+UM_AP_CLIENT_SECRET+`",
        "audience": "`+UM_AP_AUDIENCE+`"
      }`)
    req, err := http.NewRequest("POST", UM_AP_TOKEN_URL, bytes.NewBuffer(body))
    if err != nil {
        panic(err)
    }
    req.Header.Set("Content-Type", "application/json; charset=utf-8")

    client := &http.Client{}
    res, err := client.Do(req)
    if err != nil {
        panic(err)
    }
    data, err := ioutil.ReadAll(res.Body)

    var token Token
    e := json.Unmarshal([]byte(data), &token)
    if e != nil {
        panic(e)
    }
    access_token = token.Access_token
    token_type = token.Token_type
    expires_in = int64(token.Expires_in) + (time.Now().UnixNano() / int64(time.Millisecond))
    return access_token
}

func getUser(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    userId := vars["userId"]
    fmt.Println("getUser called")
    client := &http.Client{}
    url := UM_AP_API_URL+"/users/"+userId
    req, err := http.NewRequest("GET", url, nil)
    req.Header.Set("Authorization", "Bearer "+getToken())
    res, err := client.Do(req)
    if err != nil {
        panic(err)
    }
    data, err := ioutil.ReadAll(res.Body)
    fmt.Fprintln(w,string(data))
}

func getUsers(w http.ResponseWriter, r *http.Request) {
    fmt.Println("getUsers called")
    client := &http.Client{}
    url := UM_AP_API_URL+"/users"
    req, err := http.NewRequest("GET", url, nil)
    req.Header.Set("Authorization", "Bearer "+getToken())
    res, err := client.Do(req)
    if err != nil {
        panic(err)
    }
    data, err := ioutil.ReadAll(res.Body)
    fmt.Fprintln(w,string(data))
}

func getUserGroups(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    userId := vars["userId"]
    fmt.Println("getUserGroups called")
    client := &http.Client{}
    url := UM_AP_API_URL+"/users/"+userId+"/groups"
    req, err := http.NewRequest("GET", url, nil)
    req.Header.Set("Authorization", "Bearer "+getToken())
    res, err := client.Do(req)
    if err != nil {
        panic(err)
    }
    data, err := ioutil.ReadAll(res.Body)
    fmt.Fprintln(w,string(data))
}

func getToken() string{
    if(access_token != ""){
        fmt.Println("Existing token issued")
        return access_token
    }else{
        fmt.Println("New token issued")
        return generateToken()
    }
}
