package main

import (
	"image"

	"github.com/nfnt/resize"
)

func ScaleImageMaintainAspectRatio(img image.Image, maxWidth, maxHeight uint) image.Image {
	originalWidth := img.Bounds().Dx()
	originalHeight := img.Bounds().Dy()

	newWidth, newHeight := calculateNewDimensions(originalWidth, originalHeight, maxWidth, maxHeight)

	return resize.Resize(newWidth, newHeight, img, resize.NearestNeighbor)
}

func calculateNewDimensions(originalWidth, originalHeight int, maxWidth, maxHeight uint) (newWidth, newHeight uint) {
	widthRatio := float64(originalWidth) / float64(maxWidth)
	heightRatio := float64(originalHeight) / float64(maxHeight)

	if widthRatio > 1 || heightRatio > 1 {
		if widthRatio > heightRatio {
			newWidth = maxWidth
			newHeight = uint(float64(originalHeight) / widthRatio)
		} else {
			newWidth = uint(float64(originalWidth) / heightRatio)
			newHeight = maxHeight
		}
	} else {
		newWidth = uint(originalWidth)
		newHeight = uint(originalHeight)
	}

	return
}

func ContentTypeFromMimeType(mimeType string) ContentType {
	switch mimeType {
	case "image/jpeg":
		return ContentTypeJPEG
	case "image/png":
		return ContentTypePNG
	case "image/gif":
		return ContentTypeGIF
	case "image/bmp":
		return ContentTypeBMP
	default:
		return ContentTypeUnknown
	}
}

func MimeTypeFromContentType(contentType ContentType) string {
	switch contentType {
	case ContentTypeJPEG:
		return "image/jpeg"
	case ContentTypePNG:
		return "image/png"
	case ContentTypeGIF:
		return "image/gif"
	case ContentTypeBMP:
		return "image/bmp"
	default:
		return "application/octet-stream"
	}
}
