package pkg

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func GracefulShutdown(srv *http.Server, exitCh chan int) {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM, os.Kill)

	sig := <-signalChan
	log.Printf("catch signal, %+v", sig)

	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second) // 4秒后退出
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Printf("server exiting")
	close(exitCh)
}
