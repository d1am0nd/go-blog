package server

import (
    "os"
    "io"
    "fmt"
    "strconv"
    "net/http"
    "encoding/json"

    "blog2.0/go/database"

    "github.com/julienschmidt/httprouter"
)

func AllImages(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
    images, err := database.GetAllImages()
    if err != nil {
        fmt.Println(err)
        http.Error(w, "Resource not found", http.StatusNotFound)
        return
    }
    rjson, err := json.Marshal(images)
    if err != nil {
        http.Error(w, "Problem jsonifying", http.StatusBadRequest)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.Write(rjson)
}

func CreateImage(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
    imgName, err := saveImageToDir(r, "image")
    if err != nil {
        http.Error(w, err.Error(), 219)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.Write([]byte(imgName))
}

func saveImageToDir(r *http.Request, key string) (string, error){
    r.ParseMultipartForm(32 << 20)
    image, handler, err := r.FormFile(key)
    defer image.Close()
    if err != nil {
        return "", err
    }

    // Prepand number if file already exists
    i := 0
    fname := handler.Filename
    dest := "./../public/uploads/"

    for FileExists(dest + fname) {
        i = i + 1
        fname = strconv.Itoa(i) + handler.Filename
    }

    f, err := os.OpenFile(dest + fname, os.O_WRONLY|os.O_CREATE, 0666)
    defer f.Close()
    if err != nil {
        return "", err
    }
    io.Copy(f, image)
    return fname, err
}
