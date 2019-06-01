package flex

import (
	"context"

	"gitlab.fidibo.com/backend/galaxy/api/models"
)

type Publisher struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

func newPublisherByID(id int) (Publisher, error) {
	p, err := models.FindPublisherG(context.Background(), id)
	if err != nil {
		return Publisher{}, err
	}

	pub := Publisher{
		ID:    p.ID,
		Title: p.Title,
	}

	return pub, nil
}
