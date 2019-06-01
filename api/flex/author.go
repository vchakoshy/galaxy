package flex

import (
	"context"

	"gitlab.fidibo.com/backend/galaxy/api/models"
)

type Author struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func newAuthorByID(id int) (Author, error) {
	p, err := models.FindAuthorG(context.Background(), id)
	if err != nil {
		return Author{}, err
	}

	pub := Author{
		ID:   p.ID,
		Name: p.Name,
	}

	return pub, nil
}
