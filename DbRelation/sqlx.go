package main

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Post struct {
	Id         int
	Content    string
	AuthorName string `db: author`
}

var Db *sqlx.DB

func init() {
	var err error
	Db, err = sqlx.Open("postgres", "user=gwp dbname=gwp passowrd='' sslmode=disable")
	if err != nil {
		panic(err)
	}
}

func GetPost(id int) (post Post, err error) {
	post = Post{}
	err = Db.QueryRowx("SELECT id, content, author FROM posts WHERE id = $1", id).StructScan(&post)
	if err != nil {
		return
	}
	return
}

func (post *Post) Create() (err error) {
	err = Db.QueryRow("INSERT INTO posts (content, author) value ($1, $2) returning id",
		post.Content, post.AuthorName).Scan(&post.Id)
	return
}

func main() {
	post := Post{Content: "Hello World!", AuthorName: "Sau Sheong"}
	post.Create()
	fmt.Println(post)
	readPost := Post{}
	readPost, _ = GetPost(post.Id)
	fmt.Println(readPost)
}
