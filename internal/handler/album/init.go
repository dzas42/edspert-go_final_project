package handler

import (
   usecase "final-project/internal/usecase/album"
   "github.com/gin-gonic/gin"
)

type AlbumHandler interface {
   Get(context *gin.Context)
   Create(context *gin.Context)
   GetAllAlbum(context *gin.Context)
   BatchCreate(context *gin.Context)
   Update(context *gin.Context)
   Delete(context *gin.Context)
}

type albumHandler struct {
   albumUsecase usecase.AlbumUseCase
}

func NewAlbumHandler(albumUsecase usecase.AlbumUseCase) AlbumHandler {
   return &albumHandler{
      albumUsecase: albumUsecase,
   }
}