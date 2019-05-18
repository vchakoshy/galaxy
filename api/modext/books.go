package modext

import (
	"strings"

	"gitlab.fidibo.com/backend/galaxy/api/models"
)

const (
	baseImageURL = "/images/books/"
)

func getNormalImagePath(b *models.Book) string {
	return strings.Replace(b.ImageName.String, ".jpg", "_normal.jpg", 1)
}

func GetBookNormalImage(b *models.Book) string {
	return baseImageURL + getNormalImagePath(b)
}
