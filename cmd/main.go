package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/goNiki/Subscription-service/app/container"
)

var configpath = "./configs/config.yaml"

func main() {

	c, err := container.NewContainer(configpath)
	if err != nil {
		log.Fatal(err)
	}

	r := chi.NewRouter()

	r.Mount("/api", c.SubscriptionsServer)

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
