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

    err := SQL.Select(&images, "SELECT *, CONCAT('/static/uploads/', " + imageT + ".path) as path FROM " + imageT)
    return images, err
}

func FindImageById(id uint32) (Image, error) {
    image := Image{}

    err := SQL.Get(&image, "SELECT *, CONCAT('/static/uploads/', " + imageT + ".path) as path FROM " + imageT + " WHERE id = ?", id)
    return image, err
}

func FindUsersImageById(userId uint32, id uint32) (Image, error) {
    image := Image{}

    err := SQL.Get(&image, "SELECT * FROM " + imageT + " WHERE id = ? AND user_id = ?", id, userId)
    return image, err
}

func CreateImage(image *Image, userId uint32) error {
    now := time.Now()

    image.CreatedAt = timeToDb(&now)
    image.UpdatedAt = timeToDb(&now)

    stmt, err := SQL.Prepare("INSERT INTO " + imageT + " (user_id, path, name, created_at, updated_at) VALUES (?, ?, ?, ?, ?)")
    if err != nil {
        return err
    }
    _, err = stmt.Exec(userId, image.Path, image.Name, image.CreatedAt, image.UpdatedAt)
    return err
}

func UpdateImageById(image *Image, userId uint32, id uint32) error {
    now := time.Now()
    image.UpdatedAt = timeToDb(&now)

    stmt, err := SQL.Prepare("UPDATE " + imageT + " SET path=?, name=? WHERE user_id=? AND id=?")
    _, err = stmt.Exec(image.Path, image.Name, userId, id)
    return err
}