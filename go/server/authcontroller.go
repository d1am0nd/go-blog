package server

import (
    "fmt"
    "net/http"
    "encoding/json"

    "blog2.0/go/config"
    "blog2.0/go/database"

    "github.com/julienschmidt/httprouter"
    "golang.org/x/crypto/bcrypt"
)

/**
 * 1. Grabs user by param's email from db
 * 2. Compares Hases
 * 3. Returns user json (todo: jwt)
 */
func Login(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
    email := r.FormValue("email")
    pass := r.FormValue("password")

    fmt.Println(email)
    fmt.Println(pass)

    user, err := database.FindUserByEmail(email)
    if err != nil {
        http.Error(w, "Authorization failed", http.StatusUnauthorized)
        return
    }
    fmt.Println(user)

    err = compareHashAndPassword(pass, user.Password)
    if err != nil {
        http.Error(w, "Authorization failed", http.StatusUnauthorized)
        return
    }

    var rjson []byte
    rjson, err = json.Marshal(user)
    if err != nil {
        http.Error(w, "Problem jsonifying", http.StatusBadRequest)
        return
    }

    fmt.Println(rjson)

    claims := NewClaims(user.Id, config.Jwt)
    token := CreateToken(claims, config.Jwt)

    w.Header().Set("Authorization", token)

    w.Header().Set("Content-Type", "application/json")
    w.Write([]byte(rjson))
}

func CurrentUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
    w.Write([]byte("test"))
}


func compareHashAndPassword(password string, hash string) (error) {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err
}
