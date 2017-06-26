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

func UpdatePost(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
    var userId = r.Context().Value("claims").(Claims).UserId
    slug := p.ByName("slug")

    // Find existing post or 404
    post, err := database.FindOnlyMyPostBySlug(userId, slug)
    if err != nil {
        fmt.Println(err)
        http.Error(w, "Resource not found", http.StatusNotFound)
        return
    }

    // Setting values
    active := r.FormValue("active")
    if active == "1" || active == "true" {
        post.Active = true
    } else {
        post.Active = false
    }
    publishedAtValid := r.FormValue("published_at[Valid]")
    if publishedAtValid == "true" || publishedAtValid == "1" {
        post.PublishedAt.Valid = true
        post.PublishedAt.String = r.FormValue("published_at[String]")
    } else {
        post.PublishedAt.Valid = false
        post.PublishedAt.String = ""
    }
    post.Title = r.FormValue("title")
    post.Slug = r.FormValue("slug")
    post.Content = r.FormValue("content")

    err = database.UpdatePostBySlug(&post, userId, slug)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
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
