package usecase

import (
	"context"
	"final-project/internal/entity"
)

func (a artistUseCase) Get(ctx context.Context, id int64) (*entity.Artist, error) {
	artist, err := a.repo.GeArtistCache(ctx, id)
	if err == nil {
		return artist, err
	}

	artist, err = a.repo.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	// TODO: cache artist
	return artist, nil
}

func (a artistUseCase) Create(ctx context.Context, artist *entity.Artist) (*entity.Artist, error) {
	artistID, err := a.repo.Create(ctx, artist)
	artist.ID = artistID
	_ = a.repo.SetArtistCache(ctx, artistID, *artist)
	if err != nil {
		return nil, err
	}
	if err != nil {
		return nil, err
	}
	return artist, nil
}

func (a artistUseCase) GetAllArtist(ctx context.Context, limit int, page int) ([]entity.Artist, error) {
	artists, err := a.repo.GetAllArtist(ctx, limit, page)
	if err != nil {
		return nil, err
	}
	return artists, nil
}

func (a artistUseCase) BatchCreate(ctx context.Context, artists []entity.Artist) ([]int64, error) {
	ids, err := a.repo.BatchCreate(ctx, artists)
	if err != nil {
		return nil, err
	}
	return ids, nil
}

func (a artistUseCase) Update(ctx context.Context, artist entity.Artist) (*entity.Artist, error) {
	err := a.repo.Update(ctx, artist)
	_ = a.repo.SetArtistCache(ctx, artist.ID, artist)

	if err != nil {
		return nil, err
	}
	return &artist, nil
}

func (a artistUseCase) Delete(ctx context.Context, id int64) error {
	if err := a.repo.Delete(ctx, id); err != nil {
		return err
	}
	_ = a.repo.DeleteArtistCache(ctx, id)
	return nil
}
