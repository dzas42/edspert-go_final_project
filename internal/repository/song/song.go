package repository

import (
	"context"
	"encoding/json"
	"errors"
	"final-project/internal/config"
	"final-project/internal/entity"
	"fmt"
	"time"

	"gorm.io/gorm"
)

const (
	songDetailKey = "song:%d"
	expiration    = time.Hour * 1
)

func (r songRepository) Get(ctx context.Context, id int64) (*entity.Song, error) {
	var song entity.Song
	err := r.db.WithContext(ctx).Session(&gorm.Session{PrepareStmt: true}).First(&song).Where(`id = ?`, id).Error
	if err != nil {
		return nil, err
	}
	// TODO: Set cache
	return &song, nil
}

func (r songRepository) List(ctx context.Context, limit int, page int) ([]*entity.Song, error) {
	var songs []*entity.Song
	offset := (0)
	if limit == 0 {
		limit = config.DEFAULT_LIMIT
	}
	if page > 0 {
		offset = (page - 1) * limit
	}
	err := r.db.WithContext(ctx).Session(&gorm.Session{PrepareStmt: true}).Limit(limit).Offset(offset).Find(&songs).Error
	if err != nil {
		return nil, err
	}
	return songs, nil
}

func (r songRepository) Create(ctx context.Context, song *entity.Song) (*entity.Song, error) {
	err := r.db.WithContext(ctx).Session(&gorm.Session{PrepareStmt: true}).Create(&song).Error
	if err != nil {
		return nil, err
	}
	return song, nil
}

func (r songRepository) Update(ctx context.Context, song *entity.Song) (*entity.Song, error) {
	err := r.db.WithContext(ctx).Session(&gorm.Session{PrepareStmt: true}).Where(`id = ?`, song.ID).Updates(&song).Error
	if err != nil {
		return nil, err
	}
	return song, nil
}

func (r songRepository) Delete(ctx context.Context, id int64) error {
	if err := r.db.WithContext(ctx).Session(&gorm.Session{PrepareStmt: true}).Delete(entity.Song{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (r songRepository) SetSongCache(ctx context.Context, id int64, song entity.Song) error {
	songKey := fmt.Sprintf(songDetailKey, id)
	if song.ID == 0 {
		return errors.New("error")
	}
	songString, err := json.Marshal(song)
	if err != nil {
		return err
	}
	return r.cache.Set(songKey, songString, expiration).Err()
}
func (r songRepository) GeSongCache(ctx context.Context, id int64) (*entity.Song, error) {
	key := fmt.Sprintf(songDetailKey, id)
	songString, err := r.cache.Get(key).Result()

	if err != nil {
		return nil, err
	}

	var song entity.Song
	err = json.Unmarshal([]byte(songString), &song)
	if err != nil {
		return nil, err
	}
	return &song, nil
}

func (r songRepository) DeleteSongCache(ctx context.Context, id int64) error {
	key := fmt.Sprintf(songDetailKey, id)

	return r.cache.Del(key).Err()
}
