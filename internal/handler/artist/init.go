package handler

import (
	usecase "final-project/internal/usecase/artist"

	"github.com/gin-gonic/gin"
)

type ArtistHandler interface {
	Get(context *gin.Context)
	Create(context *gin.Context)
	GetAllArtist(context *gin.Context)
	Update(context *gin.Context)
	Delete(context *gin.Context)
}

type artistHandler struct {
	artistUseCase usecase.ArtistUseCase
}

func NewArtistHandler(uc usecase.ArtistUseCase) ArtistHandler {
	return &artistHandler{uc}

}
