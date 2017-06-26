package server

import (
    "fmt"
    "net/http"
    "encoding/json"

    "blog2.0/go/database"

    "github.com/julienschmidt/httprouter"
)

func PostBySlug(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
    var claims = r.Context().Value("claims")
    slug := p.ByName("slug")

    var err error
    var post database.Post

    if claims != nil {
        var userId = claims.(Claims).UserId
        post, err = database.FindMyPostBySlug(userId, slug)
    } else {
        post, err = database.FindActivePostBySlug(slug)
    }
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

func MyPosts(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
    var userId = r.Context().Value("claims").(Claims).UserId
    posts, err := database.GetUsersPosts(userId)
    if err != nil {
        fmt.Println(err)
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
