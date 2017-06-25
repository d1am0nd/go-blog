package server

import (
    "net/http"

    "github.com/julienschmidt/httprouter"
)


func NewRouter() *httprouter.Router {
    r := httprouter.New()

    r.GET("/api/posts/", ActivePosts)
    r.GET("/api/posts/:slug", PostBySlug)
    r.ServeFiles("/static/*filepath", http.Dir("../public"))

    return r
}
