package server

import (
    "net/http"

    "github.com/julienschmidt/httprouter"
)


func NewRouter() *httprouter.Router {
    r := httprouter.New()

    r.GET("/api/users/current", AuthOnly(CurrentUser))
    r.POST("/api/users/login", Login)
    r.GET("/api/posts/all", WithUser(ActivePosts))
    r.POST("/api/posts/edit/:id", AuthOnly(UpdatePost))
    r.POST("/api/posts/create", AuthOnly(CreatePost))
    r.GET("/api/posts/my/all", AuthOnly(MyPosts))
    r.GET("/api/posts/single/:slug", WithUser(PostBySlug))
    r.GET("/api/images/all", AuthOnly(AllImages))
    r.ServeFiles("/static/*filepath", http.Dir("../public"))


    /* These have to match the ones in vue router and redirect to Home */
    r.GET("/login", Home)
    r.GET("/admin/*we", Home)
    r.GET("/posts/*we", Home)
    r.GET("/", Home)
    return r
}
