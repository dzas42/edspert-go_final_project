package usecase

import (
   "context"
   "final-project/internal/entity"
   "final-project/internal/mocks"
   "github.com/stretchr/testify/mock"
   "reflect"
   "testing"
)

func Test_songUseCase_Find(t *testing.T) {
   type mockSongRepoFunc func() *mocks.SongRepostory
   expectRepo := func(f mockSongRepoFunc) *mocks.SongRepostory {
      return f()
   }
   type fields struct {
      repo *mocks.SongRepostory
   }
   type args struct {
      ctx context.Context
      id  int64
   }
   var tests = []struct {
      name    string
      fields  fields
      args    args
      want    *entity.Song
      wantErr bool
   }{
      // TODO: Add test cases.
      {
         name: "get song cache success",
         fields: fields{
            expectRepo(func() *mocks.SongRepostory {
               
               resultData := &entity.Song{
                  ID:     1,
                  Title:  "tester",
                  Lyrics: "",
               }
               repo := &mocks.SongRepostory{}
               repo.On("GeSongCache", mock.Anything, mock.Anything).Return(resultData, nil)
               return repo
            }),
         },
         args: args{
            context.TODO(),
            1,
         },
         want: &entity.Song{
            ID:     1,
            Title:  "tester",
            Lyrics: "",
         },
         wantErr: false,
      },
   }
   for _, tt := range tests {
      t.Run(tt.name, func(t *testing.T) {
         uc := songUseCase{
            repo: tt.fields.repo,
         }
         got, err := uc.Find(tt.args.ctx, tt.args.id)
         if (err != nil) != tt.wantErr {
            t.Errorf("Find() error = %v, wantErr %v", err, tt.wantErr)
            return
         }
         if !reflect.DeepEqual(got, tt.want) {
            t.Errorf("Find() got = %v, want %v", got, tt.want)
         }
      })
   }
}