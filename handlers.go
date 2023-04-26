package main

import (
	"bytes"
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/jmoiron/sqlx"
)

const (
	THUMBNAIL_RESOLUTION = 256
)

type Handler struct {
	DB *sqlx.DB
}

type ImageMetadataUpdate struct {
	ID       int    `json:"id"`
	FileName string `json:"file_name"`
	Width    int    `json:"width"`
	Height   int    `json:"height"`
	Size     int    `json:"size"`
	Tags     []Tag  `json:"tags"`
}

func (h *Handler) ServeImages(w http.ResponseWriter, r *http.Request) {
	// Parse the image ID from the URL path
	imageIDStr := r.URL.Query().Get("id")
	imageID, err := strconv.Atoi(imageIDStr)
	if err != nil {
		http.Error(w, "Invalid image ID", http.StatusBadRequest)
		return
	}

	// Retrieve the image data, content type, and tags from the database
	image := Image{}
	err = h.DB.Get(&image, "SELECT * FROM images WHERE id = ?", imageID)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Image not found", http.StatusNotFound)
		} else {
			http.Error(w, "Error retrieving image from the database", http.StatusInternalServerError)
		}
		return
	}

	var tags []Tag
	err = h.DB.Select(&tags, `
		SELECT gt.name, gt.tag_type
		FROM global_tags gt
		JOIN image_tags it ON gt.id = it.tag_id
		WHERE it.image_id = ?
	`, imageID)
	if err != nil {
		http.Error(w, "Error retrieving tags from the database", http.StatusInternalServerError)
		return
	}

	// Convert the image data and thumbnail to base64 encoded strings
	base64ImageData := base64.StdEncoding.EncodeToString(image.Data)
	base64Thumbnail := base64.StdEncoding.EncodeToString(image.Thumbnail)

	// Create a JSON object with all the image information and tags
	responseData := struct {
		ID          int         `json:"id"`
		FileName    string      `json:"file_name"`
		ImageData   string      `json:"image_data"`
		Thumbnail   string      `json:"thumbnail"`
		Width       int         `json:"width"`
		Height      int         `json:"height"`
		Size        int64       `json:"size"`
		ContentType ContentType `json:"content_type"`
		Tags        []Tag       `json:"tags"`
	}{
		ID:          image.ID,
		FileName:    image.FileName,
		ImageData:   base64ImageData,
		Thumbnail:   base64Thumbnail,
		Width:       image.Width,
		Height:      image.Height,
		Size:        image.Size,
		ContentType: image.ContentType,
		Tags:        tags,
	}

	// Marshal the JSON object
	jsonData, err := json.Marshal(responseData)
	if err != nil {
		http.Error(w, "Error marshaling JSON", http.StatusInternalServerError)
		return
	}

	// Set the content type header and serve the JSON data
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

func (h *Handler) UploadImage(w http.ResponseWriter, r *http.Request) {
	// Parse the request and get the file
	r.ParseMultipartForm(10 << 20) // 10 MB max size
	file, header, err := r.FormFile("image")
	if err != nil {
		http.Error(w, "Error retrieving the image file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Read the file into a byte slice
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		http.Error(w, "Error reading the image file", http.StatusInternalServerError)
		return
	}

	contentType := ContentTypeFromMimeType(r.Header.Get("Content-Type"))

	var img image.Image

	// Calculate the width, height, and size
	switch contentType {
	case ContentTypeBMP:

	case ContentTypeJPEG:
		img, err = jpeg.Decode(bytes.NewReader(fileBytes))
	case ContentTypePNG:
		img, err = png.Decode(bytes.NewReader(fileBytes))
	case ContentTypeGIF:
		img, err = gif.Decode(bytes.NewReader(fileBytes))
	case ContentTypeUnknown:
		img, _, err = image.Decode(bytes.NewReader(fileBytes))
	}
	if err != nil {
		http.Error(w, "Error decoding the image file", http.StatusInternalServerError)
		return
	}

	width := img.Bounds().Dx()
	height := img.Bounds().Dy()
	size := len(fileBytes)

	// Generate and store the thumbnail (not shown here, refer to previous answers)
	thumbnailBytes := ScaleImageMaintainAspectRatio(img, THUMBNAIL_RESOLUTION, THUMBNAIL_RESOLUTION)

	// Insert the image data into the database
	res, err := h.DB.Exec("INSERT INTO images (file_name, data, thumbnail, width, height, size, content_type) VALUES (?, ?, ?, ?, ?, ?, ?)",
		header.Filename, fileBytes, thumbnailBytes, width, height, size, contentType)
	if err != nil {
		http.Error(w, "Error saving the image to the database", http.StatusInternalServerError)
		return
	}

	id, _ := res.LastInsertId()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]int64{"id": id})
}

func (h *Handler) EditImageMetadata(w http.ResponseWriter, r *http.Request) {
	// Read the request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// Parse the request body
	var metadataUpdate ImageMetadataUpdate
	err = json.Unmarshal(body, &metadataUpdate)
	if err != nil {
		http.Error(w, "Error parsing request body", http.StatusBadRequest)
		return
	}

	// Start a new transaction
	tx, err := h.DB.Beginx()
	if err != nil {
		http.Error(w, "Error starting database transaction", http.StatusInternalServerError)
		return
	}
	defer tx.Rollback()

	// Update the image metadata in the database
	res, err := h.DB.Exec("UPDATE images SET width = ?, height = ?, size = ? WHERE id = ?",
		metadataUpdate.Width, metadataUpdate.Height, metadataUpdate.Size, metadataUpdate.ID)
	if err != nil {
		http.Error(w, "Error updating image metadata in the database", http.StatusInternalServerError)
		return
	}

	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		http.Error(w, "No image found with the given ID", http.StatusNotFound)
		return
	}

	// Delete existing image_tags associated with the image
	_, err = tx.Exec("DELETE FROM image_tags WHERE image_id = ?", metadataUpdate.ID)
	if err != nil {
		http.Error(w, "Error deleting existing image tags", http.StatusInternalServerError)
		return
	}
	// Insert new tags or update existing ones
	for _, tag := range metadataUpdate.Tags {
		var tagID int
		err = tx.Get(&tagID, "SELECT id FROM global_tags WHERE name = ? AND tag_type = ?", tag.Name, tag.TagType)
		if err != nil && err != sql.ErrNoRows {
			http.Error(w, "Error retrieving tag from global_tags", http.StatusInternalServerError)
			return
		}

		// If the tag doesn't exist, create it
		if err == sql.ErrNoRows {
			res, err := tx.Exec("INSERT INTO global_tags (name, tag_type) VALUES (?, ?)", tag.Name, tag.TagType)
			if err != nil {
				http.Error(w, "Error inserting new tag into global_tags", http.StatusInternalServerError)
				return
			}
			tagID64, _ := res.LastInsertId()
			tagID = int(tagID64)
		}

		// Link the tag to the image
		_, err = tx.Exec("INSERT INTO image_tags (image_id, tag_id) VALUES (?, ?)", metadataUpdate.ID, tagID)
		if err != nil {
			http.Error(w, "Error linking tag to image", http.StatusInternalServerError)
			return
		}
	}

	// Commit the transaction
	err = tx.Commit()
	if err != nil {
		http.Error(w, "Error committing transaction", http.StatusInternalServerError)
		return
	}

	// Return a success response
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) SearchImagesByTags(w http.ResponseWriter, r *http.Request) {
	// Get query parameters
	urlquery := r.URL.Query()
	tagsParam := urlquery.Get("tags")

	// Split tags and validate input
	tags := strings.Split(tagsParam, ",")
	if len(tags) == 0 {
		http.Error(w, "No tags provided", http.StatusBadRequest)
		return
	}

	var images []Image
	query := `
		SELECT i.id, i.file_name, i.thumbnail, i.width, i.height, i.size, gt.name, gt.tag_type
		FROM images i
		JOIN image_tags it ON i.id = it.image_id
		JOIN global_tags gt ON it.tag_id = gt.id
		WHERE gt.name IN (?)
		GROUP BY i.id
		HAVING COUNT(DISTINCT gt.name) = ?
		ORDER BY i.id
		LIMIT ? OFFSET ?;
	`

	limit := 20
	offset := 0

	// Prepare the query and replace the placeholder with the list of tag names
	query, args, err := sqlx.In(query, tags, len(tags), limit, offset)
	if err != nil {
		http.Error(w, "Error preparing query", http.StatusInternalServerError)
		return
	}
	query = h.DB.Rebind(query)

	// Execute the query
	err = h.DB.Select(&images, query, args...)
	if err != nil {
		http.Error(w, "Error executing query", http.StatusInternalServerError)
		return
	}

	// Return the result
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(images)
}
