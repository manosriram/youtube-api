package data

import (
	"github.com/google/uuid"
)

type Video struct {
	Id           uuid.UUID `json:"id" bson:"id"`
	Title        string    `json:"title" bson:"title"`
	Description  string    `json:"description" bson:"description"`
	PublishedAt  string    `json:"published_at" bson:"published_at"`
	ThumbnailUrl string    `json:"thumbnail_url" bson:"thumbnail_url"`
	VideoETag    string    `json:"video_etag" bson:"video_etag"`
}