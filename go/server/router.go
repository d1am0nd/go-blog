package server

import (
    "net/http"

    "github.com/julienschmidt/httprouter"
)


func NewRouter() *httprouter.Router {
    r := httprouter.New()

    r.GET("/", Home)

    r.GET("/api/posts/all", WithUser(ActivePosts))
    r.GET("/api/posts/single/:slug", PostBySlug)
    r.ServeFiles("/static/*filepath", http.Dir("../public"))

    return r
}
