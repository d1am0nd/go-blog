package config

import (
    "fmt"
    "encoding/json"
    "os"
)

var Mysql MysqlConf
var Jwt JwtConf

type MysqlConf struct {
    Username    string  `json: username`
    Password    string  `json: password`
    Hostname    string  `json: hostname`
    Port        int     `json: port`
    Name        string  `json: name`
    Parameter   string  `json: parameter`
}

type JwtConf struct {
    Secret      string `json: secret`
    Issuer      string `json: issuer`
}

func Init() {
    Mysql = GetMysqlConf()
    Jwt = GetJwtConf()
}

// MysqlConf
func (c MysqlConf) DSN() string {
    return c.Username +
            ":" +
            c.Password +
            "@tcp(" +
            c.Hostname +
            ":" +
            fmt.Sprintf("%d", c.Port) +
            ")/" +
            c.Name + c.Parameter
}

func GetMysqlConf() MysqlConf {
    file, err := os.Open("../config/database.json")
    if err != nil {
        panic(err)
    }

    decoder := json.NewDecoder(file)
    configuration := MysqlConf{}

    err = decoder.Decode(&configuration)
    if err != nil {
      panic(err)
    }

    return configuration
}


// JwtConf
func (c JwtConf) SecretToByteArray() []byte {
    return []byte(c.Secret)
}

func GetJwtConf() JwtConf {
    file, err := os.Open("../config/jwt.json")

    decoder := json.NewDecoder(file)
    configuration := JwtConf{}

    err = decoder.Decode(&configuration)
    if err != nil {
        panic(err)
    }

    return configuration
}
