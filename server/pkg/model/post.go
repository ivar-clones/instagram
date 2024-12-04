package model

import "time"

type PostDB struct {
	ID         int       `db:"id"`
	UpsertedAt time.Time `db:"upserted_at"`
	Caption    *string   `db:"caption"`
	Media      *string   `db:"media"`
	Author     string    `db:"author"`
}

type CreatePostRequest struct {
	Caption *string `json:"caption"`
	Media   *string `json:"media"`
}

type PostResponse struct {
	ID         int       `json:"id"`
	UpsertedAt time.Time `json:"upsertedAt"`
	Caption    *string   `json:"caption"`
	Media      *string   `json:"media"`
	Author     string    `json:"author"`
}

func MapFromCreatePostRequestToPostDB(source CreatePostRequest, author string) PostDB {
	return PostDB{
		Caption: source.Caption,
		Media:   source.Media,
		Author:  author,
	}
}

func MapFromPostDBToPostResponse(source PostDB) PostResponse {
	return PostResponse{
		ID:         source.ID,
		UpsertedAt: source.UpsertedAt,
		Caption:    source.Caption,
		Media:      source.Media,
		Author:     source.Author,
	}
}
