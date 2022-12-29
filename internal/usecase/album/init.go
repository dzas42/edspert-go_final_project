package usecase

import (
	"context"

	"final-project/internal/entity"
	albumRepository "final-project/internal/repository/album"
)

type AlbumUseCase interface {
	Get(ctx context.Context, id int64) (*entity.Album, error)
	Create(ctx context.Context, album *entity.Album) (*entity.Album, error)
	GetAllAlbum(ctx context.Context, limit int, page int, artist_id int64) ([]entity.Album, error)
	BatchCreate(ctx context.Context, albums []entity.Album) ([]int64, error)
	Update(ctx context.Context, album entity.Album) (*entity.Album, error)
	Delete(ctx context.Context, id int64) error
}

type albumUseCase struct {
	albumRepository albumRepository.AlbumRepository
}

// NewAlbumUseCase The function is to initialize the album use case
func NewAlbumUseCase(albumRepository albumRepository.AlbumRepository) AlbumUseCase {
	return &albumUseCase{
		albumRepository: albumRepository,
	}
}
