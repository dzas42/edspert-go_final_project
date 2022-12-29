package repository

import (
	"context"
	"final-project/internal/entity"

	"github.com/go-redis/redis"
	"gorm.io/gorm"
)

type ArtistRepository interface {
	Get(ctx context.Context, id int64) (*entity.Artist, error)
	Create(ctx context.Context, artist *entity.Artist) (int64, error)
	GetAllArtist(ctx context.Context, limit int, page int) ([]entity.Artist, error)
	BatchCreate(ctx context.Context, artists []entity.Artist) ([]int64, error)
	Update(ctx context.Context, artis entity.Artist) error
	Delete(ctx context.Context, id int64) error

	GeArtistCache(ctx context.Context, id int64) (*entity.Artist, error)
	SetArtistCache(ctx context.Context, id int64, song entity.Artist) error
	DeleteArtistCache(ctx context.Context, id int64) error
}

type artistRepository struct {
	postgres *gorm.DB
	cache    *redis.Client
}

func NewArtistRepository(db *gorm.DB, cache *redis.Client) ArtistRepository {
	return &artistRepository{
		db, cache,
	}
}
