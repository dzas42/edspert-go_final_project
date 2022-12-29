package usecase

import (
	"context"
	"final-project/internal/entity"
	repository "final-project/internal/repository/song"
)

type SongUseCase interface {
	Find(ctx context.Context, id int64) (*entity.Song, error)
	Create(ctx context.Context, album *entity.Song) (*entity.Song, error)
	FindAll(ctx context.Context, limit int, page int) ([]*entity.Song, error)
	Update(ctx context.Context, song entity.Song) (*entity.Song, error)
	Delete(ctx context.Context, id int64) error
}

type songUseCase struct {
	repo repository.SongRepostory
}

func NewSongUseCase(repo repository.SongRepostory) SongUseCase {
	return &songUseCase{repo}
}
