package server

import (
	"context"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/tejiriaustin/slabmark-api/controllers"
	"github.com/tejiriaustin/slabmark-api/repository"
	"github.com/tejiriaustin/slabmark-api/services"
)

func Start(
	ctx context.Context,
	service services.IServiceInterface,
	repo repository.IRepositoryInterface,
) {

	router := gin.New()

	log.Println("starting server...")

	controllers.AddRoutes(ctx, router, service, repo)

	go func() {
		if err := router.Run(); err != nil {
			log.Fatal("shutting down server...:", err.Error())
		}
	}()

	// graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := router; err != nil {
		log.Fatal(err)
	}
}
