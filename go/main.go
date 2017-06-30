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

    fmt.Println("ENV: " + config.Env.Env)
    if config.Env.IsProd() {
        fmt.Println("PRODUCTION")
    } else {
        fmt.Println("Not production")
    }

    http.ListenAndServe(":3000", router)
}
