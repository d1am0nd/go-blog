package database

import (
    "fmt"
    "time"
)

const imageT = "images"

type Image struct {
    Id uint32 `db:"id" json:"id"`
    UserId uint32 `db:"user_id" json:"user_id"`
    Path string `db:"path" json:"path"`
    Name string `db:"name" json:"name"`
    CreatedAt string `db:"created_at" json:"created_at"`
    UpdatedAt string `db:"updated_at" json:"updated_at"`
}

func (i *Image) GetId() string {
    r := ""
    r = fmt.Sprintf("%v", i.Id)
    return r
}

func (i *Image) GetUserId() string {
    r := ""
    r = fmt.Sprintf("%v", i.UserId)
    return r
}

func (i *Image) CreatedAtTime() time.Time {
    return parseDbTimestamp(&i.CreatedAt)
}

func (i *Image) UpdatedAtTime() time.Time {
    return parseDbTimestamp(&i.UpdatedAt)
}

func NewImage() Image {
    return Image{}
}

func (i *Image) IsEmpty() bool {
    if *i == (Image{}) {
        return true
    }
    return false
}

func GetAllImages() ([]Image, error) {
    images := []Image{}

    err := SQL.Select(&images, "SELECT * FROM " + imageT)
    return images, err
}
