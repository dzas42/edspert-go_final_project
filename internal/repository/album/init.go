package repository

import (
	"context"
	"time"

	"final-project/internal/entity"

	"github.com/go-redis/redis"
	"gorm.io/gorm"
)

type AlbumRepository interface {
	Get(ctx context.Context, id int64) (*entity.Album, error)
	Create(ctx context.Context, album *entity.Album) (int64, error)
	GetAllAlbum(ctx context.Context, limit int, page int, artist_id int64) ([]entity.Album, error)
	BatchCreate(ctx context.Context, albums []entity.Album) ([]int64, error)
	Update(ctx context.Context, album entity.Album) error
	Delete(ctx context.Context, id int64) error

	GetAlbumCache(ctx context.Context, id int64) (*entity.Album, error)
	GetAllAlbumCache(ctx context.Context) ([]entity.Album, error)
	SetAlbumCache(ctx context.Context, id int64, album entity.Album) error
	SetAllAlbumCache(ctx context.Context, albums []entity.Album) error
	DeleteAlbumCache(ctx context.Context, id int64) error
}

const (
	albumsKey      = "albums"
	albumDetailKey = "albums:%d"
	expiration     = time.Hour * 1
)

type albumRepository struct {
	postgres *gorm.DB
	cache    *redis.Client
}

// NewAlbumRepository The function is to initialize the album repository
func NewAlbumRepository(db *gorm.DB, client *redis.Client) AlbumRepository {
	return &albumRepository{
		postgres: db,
		cache:    client,
	}
}
