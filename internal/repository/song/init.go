package repository

import (
	"context"
	"final-project/internal/entity"

	"github.com/go-redis/redis"
	"gorm.io/gorm"
)

//go:generate mockery   --name=SongRepostory  --filename=song_repository.go --output=../../mocks
type SongRepostory interface {
	Get(ctx context.Context, id int64) (*entity.Song, error)
	List(ctx context.Context, limit int, page int) ([]*entity.Song, error)
	Create(ctx context.Context, song *entity.Song) (*entity.Song, error)
	Update(ctx context.Context, song *entity.Song) (*entity.Song, error)
	Delete(ctx context.Context, id int64) error

	GeSongCache(ctx context.Context, id int64) (*entity.Song, error)
	SetSongCache(ctx context.Context, id int64, song entity.Song) error
	DeleteSongCache(ctx context.Context, id int64) error
}

type songRepository struct {
	db    *gorm.DB
	cache *redis.Client
}

func NewSongRepository(db *gorm.DB, cache *redis.Client) SongRepostory {
	return &songRepository{db: db, cache: cache}
}
