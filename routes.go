package main

import (
    "net/http"
    "github.com/gorilla/mux"
)

type Route struct {
    Name        string
    Method      string
    Pattern     string
    HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {

    router := mux.NewRouter().StrictSlash(true)
    for _, route := range routes {
        router.
            Methods(route.Method).
            Path(route.Pattern).
            Name(route.Name).
            Handler(route.HandlerFunc)
    }

    return router
}

var routes = Routes{
    Route{
        "getUsers",
        "GET",
        "/users",
        getUsers,
    },
    Route{
        "getUser",
        "GET",
        "/users/{userId}",
        getUser,
    },
    Route{
        "getUserGroups",
        "GET",
        "/userGroups/{userId}",
        getUserGroups,
    },
    Route{
        "getGroups",
        "GET",
        "/groups",
        getGroups,
    },
    Route{
        "getGroupMembers",
        "GET",
        "/groupMembers/{groupId}",
        getGroupMembers,
    },
    Route{
        "getGroupRoles",
        "GET",
        "/groupRoles/{groupId}",
        getGroupRoles,
    },
    Route{
        "addGroupMember",
        "PATCH",
        "/addMembers/{groupId}/{userId}",
        addGroupMember,
    },
    Route{
        "deleteGroupMember",
        "DELETE",
        "/deleteMembers/{groupId}/{userId}",
        deleteGroupMember,
    },
    Route{
        "addGroup",
        "POST",
        "/addGroup/{groupName}/{groupDescription}",
        addGroup,
    },    
}