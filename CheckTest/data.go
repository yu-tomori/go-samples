package main

import (
	_ "github.com/lib/pq"
)

func (post *Post) fetch(id int) (err error) {
	err = post.Db.QueryRow("SELECT id, content, author FROM posts WHERE id = $1", id).
		Scan(&post.Id, &post.Content, &post.Author)
	return
}

func (post *Post) create() (err error) {
	statement := "INSERT INTO posts (content, author) values ($1, $2) returning id"
	stmt, err := post.Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()
	err = stmt.QueryRow(post.Content, post.Author).Scan(&post.Id)
	return
}

func (post *Post) update() (err error) {
	_, err = post.Db.Exec("UPDATE posts SET content = $2, author = $3 WHERE id = $1",
		post.Id, post.Content, post.Author)
	return
}

func (post *Post) delete() (err error) {
	_, err = post.Db.Exec("DELETE FROM posts WHERE id = $1", post.Id)
	return
}
