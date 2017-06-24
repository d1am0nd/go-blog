package server

import (
    "net/http"

    "github.com/julienschmidt/httprouter"
)


func NewRouter() *httprouter.Router {
    r := httprouter.New()

    r.ServeFiles("/static/*filepath", http.Dir("../public"))

    return r
}
