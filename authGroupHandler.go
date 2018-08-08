package main

import (
    "fmt"
    "net/http"
    "io/ioutil"
    "github.com/gorilla/mux"
    "bytes"
)

func getGroups(w http.ResponseWriter, r *http.Request) {
    fmt.Println("getGroups called")
    client := &http.Client{}
    url := UM_AP_API_URL+"/groups"
    req, err := http.NewRequest("GET", url, nil)
    req.Header.Set("Authorization", "Bearer "+getToken())
    res, err := client.Do(req)
    if err != nil {
        panic(err)
    }
    data, err := ioutil.ReadAll(res.Body)
    fmt.Fprintln(w,string(data))
}

func getGroupMembers(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    groupId := vars["groupId"]
    fmt.Println("getGroupMembers called")
    client := &http.Client{}
    url := UM_AP_API_URL+"/groups/"+groupId+"/members"
    req, err := http.NewRequest("GET", url, nil)
    req.Header.Set("Authorization", "Bearer "+getToken())
    res, err := client.Do(req)
    if err != nil {
        panic(err)
    }
    data, err := ioutil.ReadAll(res.Body)
    fmt.Fprintln(w,string(data))
}

func getGroupRoles(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    groupId := vars["groupId"]
    fmt.Println("getGroupRoles called")
    client := &http.Client{}
    url := UM_AP_API_URL+"/groups/"+groupId+"/roles"
    req, err := http.NewRequest("GET", url, nil)
    req.Header.Set("Authorization", "Bearer "+getToken())
    res, err := client.Do(req)
    if err != nil {
        panic(err)
    }
    data, err := ioutil.ReadAll(res.Body)
    fmt.Fprintln(w,string(data))
}

func addGroupMember(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    userId := vars["userId"]
    groupId := vars["groupId"]
    fmt.Println("addGroupMember called")
    url := UM_AP_API_URL+"/groups/"+groupId+"/members"
    
    var body = []byte(`["`+userId+`"]`)
    req, err := http.NewRequest("PATCH", url, bytes.NewBuffer(body))
    if err != nil {
        panic(err)
    }
    req.Header.Set("Authorization", "Bearer "+getToken())

    client := &http.Client{}
    res, err := client.Do(req)
    if err != nil {
        panic(err)
    }
    data, err := ioutil.ReadAll(res.Body)
    fmt.Fprintln(w,string(data))
}

func deleteGroupMember(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    userId := vars["userId"]
    groupId := vars["groupId"]
    fmt.Println("deleteGroupMember called")
    url := UM_AP_API_URL+"/groups/"+groupId+"/members"
    
    var body = []byte(`["`+userId+`"]`)
    req, err := http.NewRequest("DELETE", url, bytes.NewBuffer(body))
    if err != nil {
        panic(err)
    }
    req.Header.Set("Authorization", "Bearer "+getToken())

    client := &http.Client{}
    res, err := client.Do(req)
    if err != nil {
        panic(err)
    }
    data, err := ioutil.ReadAll(res.Body)
    fmt.Fprintln(w,string(data))
}

func addGroup(w http.ResponseWriter, r *http.Request) {
    fmt.Println("addGroup called")
    url := UM_AP_API_URL+"/groups"
    vars := mux.Vars(r)
    groupName := vars["groupName"]
    groupDescription := vars["groupDescription"]
    var body = []byte(`{"name": "`+groupName+`","description": "`+groupDescription+`"}`)
    req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
    if err != nil {
        panic(err)
    }
    req.Header.Set("Authorization", "Bearer "+getToken())

    client := &http.Client{}
    res, err := client.Do(req)
    if err != nil {
        panic(err)
    }
    data, err := ioutil.ReadAll(res.Body)
    fmt.Fprintln(w,string(data))
}
  




