package usecase

import (
	"context"
	"final-project/internal/entity"
	repository "final-project/internal/repository/artists"
)

type ArtistUseCase interface {
	Get(ctx context.Context, id int64) (*entity.Artist, error)
	Create(ctx context.Context, artist *entity.Artist) (*entity.Artist, error)
	GetAllArtist(ctx context.Context, limit int, page int) ([]entity.Artist, error)
	BatchCreate(ctx context.Context, artists []entity.Artist) ([]int64, error)
	Update(ctx context.Context, artist entity.Artist) (*entity.Artist, error)
	Delete(ctx context.Context, id int64) error
}

type artistUseCase struct {
	repo repository.ArtistRepository
}

func NewArtistUseCase(repo repository.ArtistRepository) ArtistUseCase {
	return &artistUseCase{repo: repo}
}
