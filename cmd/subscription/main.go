package main

import (
	"context"
	_ "embed"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/goNiki/Subscription-service/app/container"
	"github.com/goNiki/Subscription-service/internal/infrastructure/swagger"
)

var configpath = "./configs/config.yaml"

//go:embed openapi-bundled.yaml
var swaggerDoc []byte

func main() {

	c, err := container.NewContainer(configpath)
	if err != nil {
		log.Fatal(err)
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)

	swagger.RegisterRoutes(r, swaggerDoc)

	r.Mount("/", c.SubscriptionsServer)

	c.Server.Handler = r

	go func() {
		if err := c.Server.ListenAndServe(); err != nil {
			log.Fatalf("Ошибка запуска сервера %v", err)
		}
	}()

	fmt.Printf("Сервер запущен на %s порту \n", c.Config.Server.Port)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Сервер останавливается")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := c.Server.Shutdown(ctx); err != nil {
		log.Printf("ошибка при остановке сервера: %v", err)
	}

	log.Println("Сервер остановлен")

}
