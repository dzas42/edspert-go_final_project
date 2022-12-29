package usecase

import (
	"context"
	"final-project/internal/entity"
)

func (a albumUseCase) Get(ctx context.Context, id int64) (*entity.Album, error) {
	var album *entity.Album
	// get album from the cache
	album, err := a.albumRepository.GetAlbumCache(ctx, id)
	if err != nil {
		album, err = a.albumRepository.Get(ctx, id)
		if err != nil {
			return nil, err
		}
		err = a.albumRepository.SetAlbumCache(ctx, id, *album)
	}
	return album, nil
}

func (a albumUseCase) Create(ctx context.Context, album *entity.Album) (*entity.Album, error) {
	id, err := a.albumRepository.Create(ctx, album)
	if err != nil {
		return nil, err
	}
	album.ID = id
	// set cache album id to album
	_ = a.albumRepository.SetAlbumCache(ctx, id, *album)
	return album, nil
}

func (a albumUseCase) GetAllAlbum(ctx context.Context, limit int, page int, artist_id int64) ([]entity.Album, error) {
	return a.albumRepository.GetAllAlbum(ctx, limit, page, artist_id)
}

func (a albumUseCase) BatchCreate(ctx context.Context, albums []entity.Album) ([]int64, error) {
	return a.albumRepository.BatchCreate(ctx, albums)
}

func (a albumUseCase) Update(ctx context.Context, album entity.Album) (*entity.Album, error) {
	err := a.albumRepository.Update(ctx, album)
	_ = a.albumRepository.SetAlbumCache(ctx, album.ID, album)
	if err != nil {
		return nil, err
	}
	return &album, nil
}

func (a albumUseCase) Delete(ctx context.Context, id int64) error {
	err := a.albumRepository.Delete(ctx, id)
	if err != nil {
		return err
	}
	_ = a.albumRepository.DeleteAlbumCache(ctx, id)
	return nil
}
