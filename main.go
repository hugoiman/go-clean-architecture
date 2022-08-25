package main

import (
	"fmt"
	"go-clean-architecture/config"
	"go-clean-architecture/internal"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
)

func init() {
	if config.ReadEnv() != nil {
		os.Exit(0)
	}
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	headers := handlers.AllowedHeaders([]string{"Origin", "Accept", "Keep-Alive", "User-Agent", "If-Modified-Since", "Cache-Control", "Referer", "Authorization", "Content-Type", "X-Requested-With"})
	origins := handlers.AllowedOrigins([]string{"*"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS", "DELETE", "PUT", "HEAD"})

	internal.Setup(router)

	fmt.Println("Server running at :" + viper.GetString("server.port"))
	log.Fatal(http.ListenAndServe(":"+viper.GetString("server.port"), handlers.CORS(origins, headers, methods)(router)))
}

// just test
