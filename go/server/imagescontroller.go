package server

import (
    "fmt"
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
