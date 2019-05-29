package flex

import (
	"context"
	"strconv"
	"time"

	"github.com/patrickmn/go-cache"
	"gitlab.fidibo.com/backend/galaxy/api/models"
)

var publishersCache = cache.New(time.Minute*5, time.Minute*10)

type Publisher struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

func newPublisherByID(id int) (Publisher, error) {
	cacheKey := strconv.Itoa(id)
	if val, ex := genericsCache.Get(cacheKey); ex {
		s := val.(Publisher)
		return s, nil
	}

	p, err := models.FindPublisherG(context.Background(), id)
	if err != nil {
		return Publisher{}, err
	}

	pub := Publisher{
		ID:    p.ID,
		Title: p.Title,
	}

	publishersCache.Set(cacheKey, pub, cache.DefaultExpiration)

	return pub, nil

}
