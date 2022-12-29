package repository

import (
	"context"
	"encoding/json"
	"final-project/internal/config"
	"fmt"

	"final-project/internal/entity"

	"github.com/go-redis/redis"
	"gorm.io/gorm"
)

func (r albumRepository) Get(ctx context.Context, id int64) (*entity.Album, error) {
	var album entity.Album
	err := r.postgres.WithContext(ctx).Session(&gorm.Session{PrepareStmt: true}).Where(`id = ? `, id).Preload("Artist").First(&album).Error
	if err != nil {
		return nil, err
	}

	return &album, nil
}

func (r albumRepository) Create(ctx context.Context, album *entity.Album) (int64, error) {
	result := r.postgres.WithContext(ctx).Session(&gorm.Session{PrepareStmt: true}).Create(album)
	if result.Error != nil {
		return 0, result.Error
	}
	return album.ID, nil
}

func (r albumRepository) GetAllAlbum(ctx context.Context, limit int, page int, artist_id int64) ([]entity.Album, error) {
	var albums []entity.Album
	offset := (0)
	if limit == 0 {
		limit = config.DEFAULT_LIMIT
	}
	if page > 0 {
		offset = (page - 1) * limit
	}
	query := r.postgres.WithContext(ctx).Session(&gorm.Session{PrepareStmt: true}).Preload("Artist").Find(&albums)
	if artist_id > 0 {
		query = query.Where("artist_id =?", artist_id)
	}
	err := query.Limit(int(limit)).Offset(int(offset)).Error
	if err != nil {
		return nil, err
	}
	return albums, nil
}

func (r albumRepository) BatchCreate(ctx context.Context, albums []entity.Album) ([]int64, error) {
	err := r.postgres.WithContext(ctx).Session(&gorm.Session{PrepareStmt: true}).Create(albums).Error
	if err != nil {
		return nil, err
	}
	result := []int64{}
	for _, album := range albums {
		result = append(result, album.ID)
	}
	return result, nil
}

func (r albumRepository) Update(ctx context.Context, album entity.Album) error {
	err := r.postgres.WithContext(ctx).Session(&gorm.Session{PrepareStmt: true}).Model(album).Where(`id=?`, album.ID).Updates(album).Error
	if err != nil {
		return err
	}
	return nil
}

func (r albumRepository) Delete(ctx context.Context, id int64) error {
	return r.postgres.WithContext(ctx).Session(&gorm.Session{PrepareStmt: true}).Delete(entity.Album{}, id).Error
}

func (r albumRepository) GetAlbumCache(ctx context.Context, id int64) (*entity.Album, error) {
	var album entity.Album

	key := fmt.Sprintf(albumDetailKey, id)
	albumsString, err := r.cache.Get(key).Result()
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal([]byte(albumsString), &album)
	if err != nil {
		return &album, err
	}

	return &album, nil

}

func (r albumRepository) GetAllAlbumCache(ctx context.Context) ([]entity.Album, error) {
	var albums []entity.Album

	albumsString, err := r.cache.Get(albumsKey).Result()
	if err == redis.Nil {
		return albums, nil
	}
	if err != nil {
		return albums, err
	}

	err = json.Unmarshal([]byte(albumsString), &albums)
	if err != nil {
		return albums, err
	}

	return albums, nil
}

func (r albumRepository) SetAlbumCache(ctx context.Context, id int64, album entity.Album) error {
	var albumKey string
	albumKey = fmt.Sprintf(albumDetailKey, id)
	albumsString, err := json.Marshal(album)
	if err != nil {
		return err
	}
	return r.cache.Set(albumKey, albumsString, expiration).Err()
}

func (r albumRepository) SetAllAlbumCache(ctx context.Context, albums []entity.Album) error {
	return r.cache.Set(albumsKey, albums, expiration).Err()
}

func (r albumRepository) DeleteAlbumCache(ctx context.Context, id int64) error {
	var albumKey string
	albumKey = fmt.Sprintf(albumDetailKey, id)
	return r.cache.Del(albumsKey, albumKey).Err()

}
