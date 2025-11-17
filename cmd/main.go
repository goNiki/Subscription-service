package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/goNiki/Subscription-service/shared/pkg/openapi/subscriptions/v1"
)

var httpPort = "8082"

func main() {

	subscriptionsServer, err := subscriptions.NewServer(subscriptions.UnimplementedHandler{})
	if err != nil {
		log.Fatalf("Ошибка создания openApi сервера: %v", err)
	}

	r := chi.NewRouter()

	r.Mount("/api", subscriptionsServer)

	server := &http.Server{
		Addr:        net.JoinHostPort("localhost", httpPort),
		Handler:     r,
		ReadTimeout: 3,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Fatalf("Ошибка запуска сервера %v", err)
		}
	}()

	fmt.Printf("Сервер запущен на %s порту \n", httpPort)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Сервер останавливается")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Printf("ошибка при остановке сервера: %v", err)
	}

	log.Println("Сервер остановлен")

}
