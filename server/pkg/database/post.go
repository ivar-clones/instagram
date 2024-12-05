package database

import (
	"context"
	"instagram/pkg/model"
	"log"

	"github.com/jackc/pgx/v5"
)

type PostRepository interface {
	CreatePost(post model.PostDB) error
	GetAllPosts(username string) ([]model.PostDB, error)
	DeletePost(username string, id int) error
	GetPost(id int) (model.PostDB, error)
}

func (d *database) CreatePost(post model.PostDB) error {
	query := "insert into posts (caption, media, author) values (@caption, @media, @author)"
	args := pgx.NamedArgs{
		"caption": post.Caption,
		"media":   post.Media,
		"author":  post.Author,
	}

	if _, err := d.db.Exec(context.Background(), query, args); err != nil {
		log.Println("error inserting into posts: ", err.Error())
		return err
	}

	return nil
}

func (d *database) GetAllPosts(username string) ([]model.PostDB, error) {
	rows, err := d.db.Query(context.Background(), "select * from posts where author = $1", username)
	if err != nil {
		log.Println("error fetching posts: ", err.Error())
		return nil, err
	}
	defer rows.Close()

	posts, err := pgx.CollectRows(rows, pgx.RowToStructByName[model.PostDB])
	if err != nil {
		log.Println("error collecting rows: ", err.Error())
		return nil, err
	}

	return posts, nil
}

func (d *database) DeletePost(username string, id int) error {
	if _, err := d.db.Exec(context.Background(), "delete from posts where id = $1 and author = $2", id, username); err != nil {
		log.Println("error deleting post: ", err.Error())
		return err
	}

	return nil
}

func (d *database) GetPost(id int) (model.PostDB, error) {
	rows, err := d.db.Query(context.Background(), "select * from posts where id = $1", id)
	if err != nil {
		log.Println("error fetching posts: ", err.Error())
		return model.PostDB{}, err
	}
	defer rows.Close()

	post, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[model.PostDB])
	if err != nil {
		log.Println("error collecting row: ", err.Error())
		return model.PostDB{}, err
	}

	return post, nil
}
