package repository

import (
	"context"
	"encoding/json"
	"final-project/internal/config"
	"final-project/internal/entity"
	"fmt"
	"time"

	"gorm.io/gorm"
)

const (
	artistDetailKey = "artist:%d"
	expiration      = time.Hour * 1
)

func (r artistRepository) Get(ctx context.Context, id int64) (*entity.Artist, error) {
	var artist entity.Artist
	err := r.postgres.WithContext(ctx).Session(&gorm.Session{PrepareStmt: true}).Preload("Albums").Where(`id = ? `, id).First(&artist).Error
	if err != nil {
		return nil, err
	}

	return &artist, nil
}

func (r artistRepository) Create(ctx context.Context, artist *entity.Artist) (int64, error) {
	result := r.postgres.WithContext(ctx).Session(&gorm.Session{PrepareStmt: true}).Create(artist)
	if result.Error != nil {
		return 0, result.Error
	}
	return artist.ID, nil
}

func (r artistRepository) GetAllArtist(ctx context.Context, limit int, page int) ([]entity.Artist, error) {
	offset := 0
	if limit == 0 {
		limit = config.DEFAULT_LIMIT
	}
	if page > 0 {
		offset = (page - 1) * limit
	}
	var artists []entity.Artist
	err := r.postgres.WithContext(ctx).Session(&gorm.Session{PrepareStmt: true}).Limit(limit).Offset(offset).Find(&artists).Error
	if err != nil {
		return nil, err
	}
	return artists, nil
}

func (r artistRepository) BatchCreate(ctx context.Context, artists []entity.Artist) ([]int64, error) {
	err := r.postgres.WithContext(ctx).Session(&gorm.Session{PrepareStmt: true}).Create(artists).Error
	if err != nil {
		return nil, err
	}
	result := []int64{}
	for _, artist := range artists {
		result = append(result, artist.ID)
	}
	return result, nil
}

func (r artistRepository) Update(ctx context.Context, artist entity.Artist) error {
	err := r.postgres.WithContext(ctx).Session(&gorm.Session{PrepareStmt: true}).Model(artist).Where(`id=?`, artist.ID).Updates(artist).Error
	if err != nil {
		return err
	}
	return nil
}

func (r artistRepository) Delete(ctx context.Context, id int64) error {
	return r.postgres.WithContext(ctx).Session(&gorm.Session{PrepareStmt: true}).Delete(entity.Artist{}, id).Error

}

func (r artistRepository) GeArtistCache(ctx context.Context, id int64) (*entity.Artist, error) {
	key := fmt.Sprintf(artistDetailKey, id)

	artistString, err := r.cache.Get(key).Result()

	if err != nil {
		return nil, err
	}

	var artist entity.Artist

	err = json.Unmarshal([]byte(artistString), &artist)

	if err != nil {
		return nil, err
	}
	return &artist, nil
}

func (r artistRepository) SetArtistCache(ctx context.Context, id int64, song entity.Artist) error {
	artistKey := fmt.Sprintf(artistDetailKey, id)

	artistString, err := json.Marshal(artistKey)
	if err != nil {
		return err
	}
	return r.cache.Set(artistKey, artistString, expiration).Err()
}

func (r artistRepository) DeleteArtistCache(ctx context.Context, id int64) error {
	key := fmt.Sprintf(artistDetailKey, id)

	return r.cache.Del(key).Err()
}
