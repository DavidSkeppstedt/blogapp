package database

import (
	"database/sql"
	"github.com/DavidSkeppstedt/blog/credentialsutil"
	"github.com/DavidSkeppstedt/blog/models"
	_ "github.com/lib/pq"
	"log"
	"strconv"
)

type DbConn struct {
	db *sql.DB
}

func Open() (conn *DbConn, err error) {
	dbConf, err := credentialsutil.CreateConfig("config/db.conf")
	if err != nil {
		log.Println("Error with database config", err.Error())
		panic(err)
	}
	db, err := sql.Open("postgres", dbConf)

	if err != nil {
		log.Println("Problem with the database driver", err.Error())
		panic(err)
	}
	conn = &DbConn{db}
	return conn, err
}

func (this *DbConn) ListPosts(count string) (result []models.Post, err error) {
	limit, err := strconv.Atoi(count)
	if err != nil {
		return nil, err
	}
	rows, err := this.db.Query("SELECT * FROM posts LIMIT $1", limit)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var post models.Post
		rows.Scan(&post.Id, &post.Title, &post.Body, &post.Author,
			&post.CreatedAt, &post.UpdatedAt)
		result = append(result, post)
	}
	return result, err
}
