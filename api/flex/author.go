package flex

import (
	"context"
	"strconv"
	"time"

	"github.com/patrickmn/go-cache"
	"gitlab.fidibo.com/backend/galaxy/api/models"
)

var authorsCache = cache.New(time.Minute*5, time.Minute*10)

type Author struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func newAuthorByID(id int) (Author, error) {
	cacheKey := strconv.Itoa(id)
	if val, ex := genericsCache.Get(cacheKey); ex {
		s := val.(Author)
		return s, nil
	}

	p, err := models.FindAuthorG(context.Background(), id)
	if err != nil {
		return Author{}, err
	}

	pub := Author{
		ID:   p.ID,
		Name: p.Name,
	}

	authorsCache.Set(cacheKey, pub, cache.DefaultExpiration)

	return pub, nil

}
