package server

import (
    "fmt"
    "time"

    "gopkg.in/mgo.v2/bson"
    "github.com/dgrijalva/jwt-go"
)

type JwtConfig interface {
    GetIssuer() string
    GetSecret() string
}

type Claims struct {
    UserId bson.ObjectId `json: "user_id"`
    jwt.StandardClaims
}

func NewClaims(uid bson.ObjectId, jwtConfig JwtConfig) Claims {
    return Claims{
        uid,
        jwt.StandardClaims{
            ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
            Issuer: jwtConfig.GetIssuer() } }
}

func CreateToken(claims Claims, jwtConfig JwtConfig) string {
    claims.Issuer = jwtConfig.GetIssuer()

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

    tokenString, err := token.SignedString([]byte(jwtConfig.GetSecret()))
    if(err != nil) {
        panic(err)
    }
    return tokenString
}

func ValidateToken(myToken string, jwtConfig JwtConfig) (Claims, error) {
    claims := Claims{}

    token, err := jwt.ParseWithClaims(myToken, &claims, func(token *jwt.Token) (interface{}, error) {
        return []byte(jwtConfig.GetSecret()), nil
    })

    fmt.Println(claims.ExpiresAt)

    if err == nil && token.Valid {
        fmt.Println("Your token is valid.  I like your style.")
    } else {
        fmt.Println("This token is terrible!  I cannot accept this.")
    }

    return claims, err
}
