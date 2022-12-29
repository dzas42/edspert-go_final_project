package handler

import (
	usecase "final-project/internal/usecase/song"
	"github.com/gin-gonic/gin"
)

type SongHandler interface {
   Get(context *gin.Context)
   Create(context *gin.Context)
   GetAllSong(context *gin.Context)
   Update(context *gin.Context)
   Delete(context *gin.Context)
}

type songHandler struct {
   songUseCase usecase.SongUseCase
}

func NewSongHandler(useCase usecase.SongUseCase) SongHandler {
   return &songHandler{songUseCase: useCase}
}