package server

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/spf13/viper"
)

// StartServer starts the http server
func StartServer() {

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGKILL, syscall.SIGINT, syscall.SIGQUIT)

	port := viper.GetInt("port")
	http.HandleFunc("/crontalk/translate", translateHandler)
	http.Handle("/", http.FileServer(http.Dir("./assets")))
	go func() {
		if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
			log.Fatal("Failed to start crontalk server: ", err.Error())
		}
	}()

	time.Sleep(10 * time.Millisecond)

	fmt.Printf("Crontalk server running at http://localhost:%d\n", port)

	<-stop

	log.Println("CronTalk server gracefully shutdown")

}
