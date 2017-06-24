package database

import (
    "fmt"
    "time"
)

const postT = "posts"

type Post struct {
    Id uint32 `db:"id" json:"id"`
    Active bool `db:"active" json:"active"`
    UserId uint32 `db:"user_id" json:"user_id"`
    Title string `db:"title" json:"title"`
    Slug string `db:"slug" json:"slug"`
    Content string `db:"content" json:"content"`
    Summary string `db:"summary" json:"summary"`
    PublishedAt string `db:"published_at" json:"published_at"`
    CreatedAt string `db:"created_at" json:"created_at"`
    UpdatedAt string `db:"updated_at" json:"updated_at"`
}

func (p *Post) GetId() string {
    r := ""
    r = fmt.Sprintf("%v", p.Id)
    return r
}

func (p *Post) PublishedAtTime() time.Time {
    return parseDbTimestamp(&p.PublishedAt)
}

func (p *Post) CreatedAtTime() time.Time {
    return parseDbTimestamp(&p.CreatedAt)
}

func (p *Post) UpdatedAtTime() time.Time {
    return parseDbTimestamp(&p.UpdatedAt)
}


func NewPost() Post {
    return Post{}
}

func FindPostBySlug(slug string) (Post, error) {
    var err error
    post := Post{}

    err = SQL.Get(&post, "SELECT * FROM " + postT + " WHERE slug = ? LIMIT 1", slug)

    return post, err
}
