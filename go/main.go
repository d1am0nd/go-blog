package main

import (
    "fmt"
    "net/http"

    "blog2.0/go/server"
    "blog2.0/go/config"
    "blog2.0/go/database"
)

func main() {
    fmt.Println("Hello world!")

    config.Init()

    router := server.NewRouter()

    database.Connect(config.Mysql.DSN())

    fmt.Println("Serving on port :3000")

    http.ListenAndServe(":3000", router)
}
