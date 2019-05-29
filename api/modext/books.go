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

func IsFree(b *models.Book) bool {
	if b.Free == 1 {
		return true
	}

	return false
}

func IsRtl(b *models.Book) bool {
	return true
}
