package server

import (
    "net/http"
    "encoding/json"

    "blog2.0/go/database"

    "github.com/julienschmidt/httprouter"
)

func PostBySlug(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
    slug := p.ByName("slug")

    post, err := database.FindActivePostBySlug(slug)
    if err != nil {
        http.Error(w, "Resource not found", http.StatusNotFound)
        return
    }
    json, err := json.Marshal(post)
    if err != nil {
        http.Error(w, "Problem jsonifying", http.StatusBadRequest)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.Write(json)
}


func ActivePosts(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
    posts, err := database.GetActivePosts()
    if err != nil {
        http.Error(w, "Resource not found", http.StatusNotFound)
        return
    }
    json, err := json.Marshal(posts)
    if err != nil {
        http.Error(w, "Problem jsonifying", http.StatusBadRequest)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.Write(json)
}
