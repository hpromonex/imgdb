package main

type TagType int

const (
	Normal TagType = iota
	Meta
	Author
)

type ContentType int

const (
	ContentTypeUnknown ContentType = iota
	ContentTypeJPEG
	ContentTypePNG
	ContentTypeGIF
	ContentTypeBMP
)

type Image struct {
	ID          int         `db:"id"`
	FileName    string      `db:"file_name"`
	Data        []byte      `db:"data"`
	Thumbnail   []byte      `db:"thumbnail"`
	Width       int         `db:"width"`
	Height      int         `db:"height"`
	Size        int64       `db:"size"`
	ContentType ContentType `db:"content_type"`
}

type Tag struct {
	ID      int    `db:"id"`
	Name    string `db:"name"`
	ImageID int    `db:"image_id"`
	TagType string `db:"tag_type"`
}
