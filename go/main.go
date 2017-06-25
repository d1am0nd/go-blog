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

    router := server.NewRouter()

    database.Connect("root:@tcp(localhost:3306)/hiphop_blog")

    p, err := database.FindActivePostBySlug("test")

    fmt.Println(err)
    fmt.Println(p)

    a, er := database.GetActivePosts()

    fmt.Println(er)
    fmt.Println(a)

    config.Init()

    fmt.Println(config.Mysql.DSN())

    fmt.Println(config.Jwt)

    http.ListenAndServe(":3000", router)
}
