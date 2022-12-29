package main

import (
	artistHandler "final-project/internal/handler/artist"
	songHandler "final-project/internal/handler/song"
	artistRepository "final-project/internal/repository/artists"
	songRepository "final-project/internal/repository/song"
	artistUsecase "final-project/internal/usecase/artist"

	songUsecase "final-project/internal/usecase/song"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"

	"final-project/database/postgres"
	"final-project/database/redis"
	"final-project/internal/config"
	"final-project/internal/entity"
	albumHandler "final-project/internal/handler/album"
	albumRepository "final-project/internal/repository/album"
	albumUsecase "final-project/internal/usecase/album"
)

func main() {
	conf, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("🚀 Could not load environment variables", err)
	}

	if err = postgres.InitialConnection(conf); err != nil {
		log.Fatal("🚀 Could not load database postgres", err)
	}

	db := postgres.GetConnection()

	cache := redis.InitConnection(conf)
	albumRepo := albumRepository.NewAlbumRepository(db, cache)
	albumUc := albumUsecase.NewAlbumUseCase(albumRepo)
	albumHd := albumHandler.NewAlbumHandler(albumUc)

	songRepo := songRepository.NewSongRepository(db, cache)
	songUc := songUsecase.NewSongUseCase(songRepo)
	songHd := songHandler.NewSongHandler(songUc)

	artistRepo := artistRepository.NewArtistRepository(db, cache)
	artistUc := artistUsecase.NewArtistUseCase(artistRepo)
	artistHd := artistHandler.NewArtistHandler(artistUc)

	//  check migration
	var migrate = false
	flag.BoolVar(&migrate, "migrate", migrate, "migrate table database")
	flag.Parse()
	if migrate == true {
		err := db.AutoMigrate(entity.Artist{}, entity.Album{}, entity.Song{})

		if err != nil {
			log.Fatalf("error migrate")
		}
	}

	// Initialize gin
	r := gin.Default()

	// setup routes
	albumRoutes := r.Group("/api/v1/albums")
	{
		albumRoutes.GET("/", albumHd.GetAllAlbum)
		albumRoutes.POST("/", albumHd.Create)
		albumRoutes.POST("/batch", albumHd.BatchCreate)
		albumRoutes.GET("/:id", albumHd.Get)
		albumRoutes.PUT("/:id", albumHd.Update)
		albumRoutes.DELETE("/:id", albumHd.Delete)
	}

	songRoutes := r.Group("/api/v1/songs")
	{
		songRoutes.GET("/", songHd.GetAllSong)
		songRoutes.POST("/", songHd.Create)
		songRoutes.GET("/:id", songHd.Get)
		songRoutes.PUT("/:id", songHd.Update)
		songRoutes.DELETE("/:id", songHd.Delete)
	}

	artistRoutes := r.Group("/api/v1/artists")
	{
		artistRoutes.GET("/", artistHd.GetAllArtist)
		artistRoutes.POST("/", artistHd.Create)
		artistRoutes.GET("/:id", artistHd.Get)
		artistRoutes.PUT("/:id", artistHd.Update)
		artistRoutes.DELETE("/:id", artistHd.Delete)
	}

	// Run the gin gonic specific port
	runWithPort := fmt.Sprintf("0.0.0.0:%s", conf.ServerPort)
	err = r.Run(runWithPort)
	if err != nil {
		panic(fmt.Sprintf("Can not serve gin at %s", runWithPort))
	}

}
