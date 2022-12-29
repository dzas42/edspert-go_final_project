// Code generated by mockery v2.15.0. DO NOT EDIT.

package mocks

import (
	context "context"
	entity "final-project/internal/entity"

	mock "github.com/stretchr/testify/mock"
)

// SongRepostory is an autogenerated mock type for the SongRepostory type
type SongRepostory struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, song
func (_m *SongRepostory) Create(ctx context.Context, song *entity.Song) (*entity.Song, error) {
	ret := _m.Called(ctx, song)

	var r0 *entity.Song
	if rf, ok := ret.Get(0).(func(context.Context, *entity.Song) *entity.Song); ok {
		r0 = rf(ctx, song)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Song)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *entity.Song) error); ok {
		r1 = rf(ctx, song)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: ctx, id
func (_m *SongRepostory) Delete(ctx context.Context, id int64) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteSongCache provides a mock function with given fields: ctx, id
func (_m *SongRepostory) DeleteSongCache(ctx context.Context, id int64) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GeSongCache provides a mock function with given fields: ctx, id
func (_m *SongRepostory) GeSongCache(ctx context.Context, id int64) (*entity.Song, error) {
	ret := _m.Called(ctx, id)

	var r0 *entity.Song
	if rf, ok := ret.Get(0).(func(context.Context, int64) *entity.Song); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Song)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Get provides a mock function with given fields: ctx, id
func (_m *SongRepostory) Get(ctx context.Context, id int64) (*entity.Song, error) {
	ret := _m.Called(ctx, id)

	var r0 *entity.Song
	if rf, ok := ret.Get(0).(func(context.Context, int64) *entity.Song); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Song)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// List provides a mock function with given fields: ctx, limit, page
func (_m *SongRepostory) List(ctx context.Context, limit int, page int) ([]*entity.Song, error) {
	ret := _m.Called(ctx, limit, page)

	var r0 []*entity.Song
	if rf, ok := ret.Get(0).(func(context.Context, int, int) []*entity.Song); ok {
		r0 = rf(ctx, limit, page)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*entity.Song)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int, int) error); ok {
		r1 = rf(ctx, limit, page)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetSongCache provides a mock function with given fields: ctx, id, song
func (_m *SongRepostory) SetSongCache(ctx context.Context, id int64, song entity.Song) error {
	ret := _m.Called(ctx, id, song)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, entity.Song) error); ok {
		r0 = rf(ctx, id, song)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Update provides a mock function with given fields: ctx, song
func (_m *SongRepostory) Update(ctx context.Context, song *entity.Song) (*entity.Song, error) {
	ret := _m.Called(ctx, song)

	var r0 *entity.Song
	if rf, ok := ret.Get(0).(func(context.Context, *entity.Song) *entity.Song); ok {
		r0 = rf(ctx, song)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Song)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *entity.Song) error); ok {
		r1 = rf(ctx, song)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewSongRepostory interface {
	mock.TestingT
	Cleanup(func())
}

// NewSongRepostory creates a new instance of SongRepostory. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewSongRepostory(t mockConstructorTestingTNewSongRepostory) *SongRepostory {
	mock := &SongRepostory{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
