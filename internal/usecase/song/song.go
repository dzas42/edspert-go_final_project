package usecase

import (
	"context"
	"final-project/internal/entity"
)

func (uc songUseCase) Find(ctx context.Context, id int64) (*entity.Song, error) {
	song, err := uc.repo.GeSongCache(ctx, id)
	if err == nil {
		return song, err
	}

	// Get from db
	song, err = uc.repo.Get(ctx, id)
	if err != nil {
		return song, err
	}
	_ = uc.repo.SetSongCache(ctx, id, *song)

	return song, nil
}

func (uc songUseCase) Create(ctx context.Context, song *entity.Song) (*entity.Song, error) {
	result, err := uc.repo.Create(ctx, song)
	if err != nil {
		return nil, err
	}
	song.ID = result.ID
	return song, err
}

func (uc songUseCase) FindAll(ctx context.Context, limit int, page int) ([]*entity.Song, error) {
	songs, err := uc.repo.List(ctx, limit, page)
	if err != nil {
		return nil, err
	}
	return songs, nil
}

func (uc songUseCase) Update(ctx context.Context, song entity.Song) (*entity.Song, error) {
	result, err := uc.repo.Update(ctx, &song)
	if err != nil {
		return nil, err
	}
	song.ID = result.ID
	return &song, err
}

func (uc songUseCase) Delete(ctx context.Context, id int64) error {
	return uc.repo.Delete(ctx, id)
}
