package server

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/viper"
)

// StartServer starts the http server
func StartServer() {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGKILL, syscall.SIGINT, syscall.SIGQUIT)

	port := viper.GetInt("port")
	mux := http.NewServeMux()
	mux.HandleFunc("/crontalk/translate", translateHandler)
	mux.HandleFunc("/", templateHandler)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: mux,
	}

	go func() {
		sig := <-stop

		log.Printf("CronTalk server gracefully shutting down (reason: %v)\n", sig)

		err := srv.Close()
		if err != nil {
			log.Fatalf("Error shutting down server: %v", err)
		}
	}()

	fmt.Printf("Crontalk server running at http://localhost:%d\n", port)
	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatalf("Failed to start crontalk server: %v", err)
	}

}
