package main

import (
    "fmt"
    "net/http"

    "blog2.0/go/server"
    "blog2.0/go/database"
)

func main() {
    fmt.Println("Hello world!")

    router := server.NewRouter()

    database.Connect("hidden")

    p, err := database.FindPostBySlug("test")

    fmt.Println(err)
    fmt.Println(p)

    http.ListenAndServe(":3000", router)
}
